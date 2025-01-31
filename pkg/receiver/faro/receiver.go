// SPDX-License-Identifier: Apache-2.0

package faroreceiver // import "github.com/grafana/faro/pkg/receiver/faro"

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	faro "github.com/grafana/faro/pkg/go"
	httpHelper "github.com/grafana/faro/pkg/receiver/faro/internal/httphelper"
	farotranslator "github.com/grafana/faro/pkg/translator/faro"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/component/componentstatus"
	"go.opentelemetry.io/collector/config/confighttp"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/receiver/receiverhelper"
	"go.uber.org/zap"
	spb "google.golang.org/genproto/googleapis/rpc/status"
)

func newFaroReceiver(cfg *Config, set *receiver.Settings) (*faroReceiver, error) {
	r := &faroReceiver{
		cfg:      cfg,
		settings: set,
	}

	var err error
	r.obsrepHTTP, err = receiverhelper.NewObsReport(receiverhelper.ObsReportSettings{
		ReceiverID:             set.ID,
		Transport:              "http",
		ReceiverCreateSettings: *set,
	})
	if err != nil {
		return nil, err
	}

	return r, nil
}

type faroReceiver struct {
	cfg        *Config
	serverHTTP *http.Server

	nextTraces consumer.Traces
	nextLogs   consumer.Logs

	obsrepHTTP *receiverhelper.ObsReport

	settings *receiver.Settings
}

func (r *faroReceiver) Start(ctx context.Context, host component.Host) error {
	if r.nextTraces == nil {
		return fmt.Errorf("traces consumer is not registered")
	}
	if r.nextLogs == nil {
		return fmt.Errorf("logs consumer is not registered")
	}

	return r.startHTTPServer(ctx, host)
}

func (r *faroReceiver) Shutdown(ctx context.Context) error {
	return r.serverHTTP.Shutdown(ctx)
}

func (r *faroReceiver) RegisterTracesConsumer(tc consumer.Traces) {
	r.nextTraces = tc
}

func (r *faroReceiver) RegisterLogsConsumer(lc consumer.Logs) {
	r.nextLogs = lc
}

func (r *faroReceiver) startHTTPServer(ctx context.Context, host component.Host) error {
	if r.cfg.HTTP == nil {
		return nil
	}

	httpMux := http.NewServeMux()
	httpMux.HandleFunc(r.cfg.HTTP.FaroURLPath, func(resp http.ResponseWriter, req *http.Request) {
		r.handleFaroRequest(resp, req)
	})

	var err error
	if r.serverHTTP, err = r.cfg.HTTP.ToServer(
		ctx,
		host,
		r.settings.TelemetrySettings,
		httpMux,
		confighttp.WithErrorHandler(errorHandler),
	); err != nil {
		return err
	}

	go func() {
		if err := r.serverHTTP.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) && err != nil {
			componentstatus.ReportStatus(host, componentstatus.NewFatalErrorEvent(err))
		}
	}()
	r.settings.Logger.Info("HTTP server started", zap.String("address", r.serverHTTP.Addr))

	return nil
}

func (r *faroReceiver) handleFaroRequest(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	resp.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// Preflight request
	if req.Method == http.MethodOptions {
		resp.WriteHeader(http.StatusOK)
		return
	}

	if req.Method != http.MethodPost {
		errorHandler(resp, req, "only POST method is supported", http.StatusMethodNotAllowed)
		return
	}

	if req.Header.Get("Content-Type") != "application/json" {
		errorHandler(resp, req, "invalid Content-Type", http.StatusUnsupportedMediaType)
		return
	}

	body, err := io.ReadAll(req.Body)
	if err != nil {
		errorHandler(resp, req, fmt.Sprintf("failed to read request body: %v", err), http.StatusBadRequest)
		return
	}
	defer req.Body.Close()

	var payload faro.Payload
	if err := json.Unmarshal(body, &payload); err != nil {
		errorHandler(resp, req, fmt.Sprintf("failed to parse Faro payload: %v", err), http.StatusBadRequest)
		return
	}

	var errors []string

	traces, err := farotranslator.TranslateToTraces(req.Context(), payload)
	if err == nil {
		err = r.nextTraces.ConsumeTraces(req.Context(), *traces)
		if err != nil {
			errors = append(errors, fmt.Sprintf("failed to push traces: %v", err))
		}
	} else {
		errors = append(errors, fmt.Sprintf("failed to convert Faro traces: %v", err))
	}

	logs, err := farotranslator.TranslateToLogs(req.Context(), payload)
	if err == nil {
		err = r.nextLogs.ConsumeLogs(req.Context(), *logs)
		if err != nil {
			errors = append(errors, fmt.Sprintf("failed to push logs: %v", err))
		}
	} else {
		errors = append(errors, fmt.Sprintf("failed to convert Faro logs: %v", err))
	}

	if len(errors) > 0 {
		errorHandler(resp, req, strings.Join(errors, "\n"), http.StatusInternalServerError)
		return
	}

	resp.WriteHeader(http.StatusOK)
}

func errorHandler(w http.ResponseWriter, r *http.Request, errMsg string, statusCode int) {
	s := httpHelper.NewStatusFromMsgAndHTTPCode(errMsg, statusCode)
	writeStatusResponse(w, jsEncoder, statusCode, s.Proto())
}

func writeStatusResponse(w http.ResponseWriter, enc encoder, statusCode int, rsp *spb.Status) {
	msg, err := enc.marshalStatus(rsp)
	if err != nil {
		return
	}
	writeResponse(w, enc.contentType(), statusCode, msg)
}

func writeResponse(w http.ResponseWriter, contentType string, statusCode int, msg []byte) {
	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(statusCode)
	// Nothing we can do with the error if we cannot write to the response.
	_, _ = w.Write(msg)
}
