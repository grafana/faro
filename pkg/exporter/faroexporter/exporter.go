// SPDX-License-Identifier: Apache-2.0

package faroexporter // import "github.com/grafana/faro/pkg/exporter/faroexporter"

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"runtime"
	"strconv"
	"sync"
	"time"

	httpHelper "github.com/grafana/faro/pkg/exporter/faroexporter/internal/httphelper"
	faro "github.com/grafana/faro/pkg/go"
	farotranslator "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/faro"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/consumer/consumererror"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/exporter/exporterhelper"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.uber.org/multierr"
	"go.uber.org/zap"
)

type faroExporter struct {
	config    *Config
	client    *http.Client
	logger    *zap.Logger
	settings  component.TelemetrySettings
	userAgent string

	wg sync.WaitGroup
}

const (
	headerRetryAfter         = "Retry-After"
	maxHTTPResponseReadBytes = 64 * 1024
	jsonContentType          = "application/json"
)

func newExporter(cfg component.Config, set exporter.Settings) (*faroExporter, error) {
	oCfg := cfg.(*Config)

	if oCfg.Endpoint != "" {
		_, err := url.Parse(oCfg.Endpoint)
		if err != nil {
			return nil, errors.New("endpoint must be a valid URL")
		}
	}

	userAgent := fmt.Sprintf("%s/%s (%s/%s)",
		set.BuildInfo.Description, set.BuildInfo.Version, runtime.GOOS, runtime.GOARCH)

	return &faroExporter{
		config:    oCfg,
		logger:    set.Logger,
		userAgent: userAgent,
		settings:  set.TelemetrySettings,
	}, nil
}

func (fe *faroExporter) start(ctx context.Context, host component.Host) error {
	client, err := fe.config.ClientConfig.ToClient(ctx, host, fe.settings)
	if err != nil {
		return err
	}
	fe.client = client
	return nil
}

func (fe *faroExporter) export(ctx context.Context, fp *faro.Payload) error {
	fe.logger.Debug("Preparing to make HTTP request", zap.String("endpoint", fe.config.Endpoint))
	request, err := json.Marshal(fp)
	if err != nil {
		return consumererror.NewPermanent(err)
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, fe.config.Endpoint, bytes.NewReader(request))
	if err != nil {
		return consumererror.NewPermanent(err)
	}
	req.Header.Set("Content-Type", jsonContentType)
	req.Header.Set("User-Agent", fe.userAgent)

	resp, err := fe.client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to make an HTTP request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusAccepted {
		return nil
	}

	var errString string
	var formattedErr error
	if resp.StatusCode != 0 {
		errString = fmt.Sprintf(
			"error exporting items, request to %s responded with HTTP Status Code %d, Message=%s",
			fe.config.Endpoint, resp.StatusCode, resp.Body)
	} else {
		errString = fmt.Sprintf(
			"error exporting items, request to %s responded with HTTP Status Code %d",
			fe.config.Endpoint, resp.StatusCode)
	}
	formattedErr = httpHelper.NewStatusFromMsgAndHTTPCode(errString, resp.StatusCode).Err()

	if httpHelper.IsRetryableStatusCode(resp.StatusCode) {
		retryAfter := 0
		isThrottleError := resp.StatusCode == http.StatusTooManyRequests || resp.StatusCode == http.StatusServiceUnavailable
		if val := resp.Header.Get(headerRetryAfter); isThrottleError && val != "" {
			if seconds, err := strconv.Atoi(val); err == nil {
				retryAfter = seconds
			}
		}
		return exporterhelper.NewThrottleRetry(formattedErr, time.Duration(retryAfter)*time.Second)
	}

	return consumererror.NewPermanent(formattedErr)
}

func (fe *faroExporter) consume(ctx context.Context, fp []faro.Payload) error {
	var errs error
	fe.wg.Add(len(fp))
	for _, p := range fp {
		go func() {
			defer fe.wg.Done()
			err := fe.export(ctx, &p)
			errs = multierr.Append(errs, err)
		}()
	}
	fe.wg.Wait()
	return errs
}

func (fe *faroExporter) ConsumeTraces(ctx context.Context, td ptrace.Traces) error {
	fp, err := farotranslator.TranslateFromTraces(ctx, td)
	if err != nil {
		return fmt.Errorf("failed to translate traces to faro payloads: %w", err)
	}
	return fe.consume(ctx, fp)
}

func (fe *faroExporter) ConsumeLogs(ctx context.Context, ld plog.Logs) error {
	fp, err := farotranslator.TranslateFromLogs(ctx, ld)
	if err != nil {
		return fmt.Errorf("failed to translate logs to faro payloads: %w", err)
	}
	return fe.consume(ctx, fp)
}

func (fe *faroExporter) Capabilities() consumer.Capabilities {
	return consumer.Capabilities{MutatesData: false}
}
