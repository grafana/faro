// SPDX-License-Identifier: Apache-2.0

package faro // import "github.com/grafana/faro/pkg/translator/faro"

import (
	"context"

	faro "github.com/grafana/faro/pkg/go"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/ptrace"
)

// Translator is a processor that converts OTLP data to Faro format.
// Added here as scaffolding pending future development.

type Translator interface {
	TranslateLogs(ctx context.Context, logs plog.Logs) ([]faro.Log, error)
	TranslateTraces(ctx context.Context, traces ptrace.Traces) (faro.Traces, error)
}

type FaroTranslator struct{}

func (t *FaroTranslator) TranslateLogs(_ context.Context, _ plog.Logs) ([]faro.Log, error) {
	return []faro.Log{}, nil
}

func (t *FaroTranslator) TranslateTraces(_ context.Context, _ ptrace.Traces) (faro.Traces, error) {
	return faro.Traces{}, nil
}
