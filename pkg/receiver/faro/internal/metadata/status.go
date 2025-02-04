package metadata

import (
	"go.opentelemetry.io/collector/component"
)

var (
	Type      = component.MustNewType("faroreceiver")
	ScopeName = "github.com/grafana/faro/pkg/receiver/faro"
)

const (
	TracesStability  = component.StabilityLevelDevelopment
	MetricsStability = component.StabilityLevelDevelopment
	LogsStability    = component.StabilityLevelDevelopment
)
