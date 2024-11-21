// SPDX-License-Identifier: Apache-2.0

package faroexporter // import "github.com/grafana/faro/exporter"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/exporter/exporterhelper"

	"github.com/grafana/faro/exporter/internal/metadata"
)

// NewFactory creates a factory for Faro exporter.
func NewFactory() exporter.Factory {
	return exporter.NewFactory(
		metadata.Type,
		createDefaultConfig,
		exporter.WithTraces(createTracesExporter, metadata.TracesStability),
		exporter.WithLogs(createLogsExporter, metadata.LogsStability),
		exporter.WithMetrics(createMetricsExporter, metadata.MetricsStability),
	)
}

func createDefaultConfig() component.Config {
	return &Config{
		FaroExporter: FaroExporterConfig{
			Endpoint: "",
		},
	}
}

func createLogsExporter(ctx context.Context, params exporter.Settings, config component.Config) (exporter.Logs, error) {
	faroExporter := newFaroExporter(config.(*Config), params)

	return exporterhelper.NewLogs(ctx, params, config, faroExporter.ConsumeLogs, exporterhelper.WithStart(faroExporter.start))
}

func createMetricsExporter(ctx context.Context, params exporter.Settings, config component.Config) (exporter.Metrics, error) {
	faroExporter := newFaroExporter(config.(*Config), params)

	return exporterhelper.NewMetrics(ctx, params, config, faroExporter.ConsumeMetrics, exporterhelper.WithStart(faroExporter.start))
}

func createTracesExporter(ctx context.Context, params exporter.Settings, config component.Config) (exporter.Traces, error) {
	faroExporter := newFaroExporter(config.(*Config), params)

	return exporterhelper.NewTraces(ctx, params, config, faroExporter.ConsumeTraces, exporterhelper.WithStart(faroExporter.start))
}
