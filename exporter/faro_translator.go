// SPDX-License-Identifier: Apache-2.0

package faroexporter // import "github.com/grafana/faro/exporter"

import (
	"context"

	faro "github.com/grafana/faro/pkg/go"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/ptrace"
)

// Translator is a processor that converts OTLP data to Faro format.
// Added here as scaffolding pending future development.

type Translator interface {
	TranslateMetrics(ctx context.Context, md pmetric.Metrics) ([]faro.Measurement, error)
	TranslateLogs(ctx context.Context, logs plog.Logs) ([]faro.Log, error)
	TranslateTraces(ctx context.Context, traces ptrace.Traces) (faro.Traces, error)
}

type faroTranslator struct{}

func (t *faroTranslator) TranslateMetrics(_ context.Context, _ pmetric.Metrics) ([]faro.Measurement, error) {
	return []faro.Measurement{}, nil
}

func (t *faroTranslator) TranslateLogs(_ context.Context, _ plog.Logs) ([]faro.Log, error) {
	return []faro.Log{}, nil
}

func (t *faroTranslator) TranslateTraces(_ context.Context, _ ptrace.Traces) (faro.Traces, error) {
	return faro.Traces{}, nil
}
