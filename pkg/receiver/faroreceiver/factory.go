// SPDX-License-Identifier: Apache-2.0

package faroreceiver // import "github.com/grafana/faro/pkg/receiver/faroreceiver"

import (
	"context"
	"fmt"

	"github.com/grafana/faro/pkg/receiver/faroreceiver/internal/metadata"
	"github.com/grafana/faro/pkg/receiver/faroreceiver/internal/sharedcomponent"
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

	receiver, err := receivers.LoadOrStore(
		fCfg,
		func() (*faroReceiver, error) {
			return newFaroReceiver(fCfg, &set)
		},
	)
	if err != nil {
		return nil, err
	}

	receiver.Unwrap().RegisterTracesConsumer(nextTraces)

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
	receiver, err := receivers.LoadOrStore(
		fCfg,
		func() (*faroReceiver, error) {
			return newFaroReceiver(fCfg, &set)
		},
	)
	if err != nil {
		return nil, err
	}

	receiver.Unwrap().RegisterLogsConsumer(nextLogs)

	return receiver, nil
}

// This is the map of already created Faro receivers for particular configurations.
// We maintain this map because the receiver.Factory is asked trace and metric receivers separately
// when it gets createFaroReceiverTraces() and createFaroReceiverLogs() but they must not
// create separate objects, they must use one faroReceiver object per configuration.
// When the receiver is shutdown it should be removed from this map so the same configuration
// can be recreated successfully.
var receivers = sharedcomponent.NewMap[*Config, *faroReceiver]()
