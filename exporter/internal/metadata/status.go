package metadata

import (
	"go.opentelemetry.io/collector/component"
)

var (
	Type      = component.MustNewType("faroexporter")
	ScopeName = "github.com/grafana/faro/exporter"
)

const (
	TracesStability  = component.StabilityLevelDevelopment
	MetricsStability = component.StabilityLevelDevelopment
	LogsStability    = component.StabilityLevelDevelopment
)
