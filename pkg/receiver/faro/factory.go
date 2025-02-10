// SPDX-License-Identifier: Apache-2.0

package faroreceiver // import "github.com/grafana/faro/pkg/receiver/faro"

import (
	"context"
	"fmt"

	"github.com/grafana/faro/pkg/receiver/faro/internal/metadata"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/confighttp"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
)

const (
	defaultFaroEndpoint = "localhost:8080"
)

func createDefaultConfig() component.Config {
	return &Config{
		ServerConfig: &confighttp.ServerConfig{
			Endpoint: defaultFaroEndpoint,
		},
	}
}

func NewFactory() receiver.Factory {
	return receiver.NewFactory(
		metadata.Type,
		createDefaultConfig,
		receiver.WithTraces(createFaroReceiverTraces, metadata.TracesStability),
		receiver.WithLogs(createFaroReceiverLogs, metadata.LogsStability))
}

func createFaroReceiverTraces(
	ctx context.Context,
	set receiver.Settings,
	cfg component.Config,
	nextTraces consumer.Traces,
) (receiver.Traces, error) {
	fCfg, ok := cfg.(*Config)
	if !ok {
		return nil, fmt.Errorf("invalid configuration: %T", cfg)
	}

	receiver, err := newFaroReceiver(fCfg, &set)
	if err != nil {
		return nil, err
	}

	receiver.RegisterTracesConsumer(nextTraces)

	return receiver, nil
}

func createFaroReceiverLogs(
	ctx context.Context,
	set receiver.Settings,
	cfg component.Config,
	nextLogs consumer.Logs,
) (receiver.Logs, error) {
	fCfg, ok := cfg.(*Config)
	if !ok {
		return nil, fmt.Errorf("invalid configuration: %T", cfg)
	}

	receiver, err := newFaroReceiver(fCfg, &set)
	if err != nil {
		return nil, err
	}

	receiver.RegisterLogsConsumer(nextLogs)

	return receiver, nil
}
