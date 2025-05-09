// SPDX-License-Identifier: Apache-2.0

package faroreceiver // import "github.com/grafana/faro/pkg/receiver/faroreceiver"

import (
	"errors"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/confighttp"
	"go.opentelemetry.io/collector/confmap"
)

type Config struct {
	*confighttp.ServerConfig `mapstructure:",squash"`
}

var (
	_ component.Config    = (*Config)(nil)
	_ confmap.Unmarshaler = (*Config)(nil)
)

func (cfg *Config) Validate() error {
	if cfg.Endpoint == "" {
		return errors.New("must specify endpoint")
	}
	return nil
}

func (cfg *Config) Unmarshal(conf *confmap.Conf) error {
	err := conf.Unmarshal(cfg)
	if err != nil {
		return err
	}

	return nil
}
