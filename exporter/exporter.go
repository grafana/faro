// SPDX-License-Identifier: Apache-2.0

package faroexporter // import "github.com/grafana/faro/exporter"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.uber.org/zap"
)

type faroExporter struct {
	config     *Config
	translator Translator
	client     FaroClient
	logger     *zap.Logger
}

func newFaroExporter(config *Config, params exporter.Settings) *faroExporter {
	faroExporter := &faroExporter{
		config:     config,
		translator: &faroTranslator{},
		logger:     params.Logger,
	}
	return faroExporter
}

func (fe *faroExporter) start(_ context.Context, host component.Host) error {
	return nil
}

func (fe *faroExporter) Capabilities() consumer.Capabilities {
	return consumer.Capabilities{MutatesData: false}
}

func (fe *faroExporter) ConsumeMetrics(ctx context.Context, md pmetric.Metrics) error {
	fm, err := fe.translator.TranslateMetrics(ctx, md)
	if err != nil {
		return err
	}
	return fe.client.SendMeasurements(fm)
}

func (fe *faroExporter) ConsumeLogs(ctx context.Context, logs plog.Logs) error {
	flogs, err := fe.translator.TranslateLogs(ctx, logs)
	if err != nil {
		return err
	}
	return fe.client.SendLogs(flogs)
}

func (fe *faroExporter) ConsumeTraces(ctx context.Context, traces ptrace.Traces) error {
	ftraces, err := fe.translator.TranslateTraces(ctx, traces)
	if err != nil {
		return err
	}
	return fe.client.SendTraces(ftraces)
}
