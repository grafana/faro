// SPDX-License-Identifier: Apache-2.0

package faroreceiver // import "github.com/grafana/faro/pkg/receiver/faro"

import (
	"errors"
	"fmt"
	"net/url"
	"path"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/confighttp"
	"go.opentelemetry.io/collector/confmap"
)

const protoHTTP = "protocols::http"

type HTTPConfig struct {
	*confighttp.ServerConfig `mapstructure:",squash"`
}

type Protocols struct {
	HTTP *HTTPConfig `mapstructure:"http"`
}

type Config struct {
	Protocols `mapstructure:"protocols,squash"`
}

var (
	_ component.Config    = (*Config)(nil)
	_ confmap.Unmarshaler = (*Config)(nil)
)

func (cfg *Config) Validate() error {
	if cfg.HTTP == nil {
		return errors.New("must specify HTTP protocol")
	}
	if cfg.HTTP.Endpoint == "" {
		return errors.New("must specify endpoint")
	}
	return nil
}

func (cfg *Config) Unmarshal(conf *confmap.Conf) error {
	err := conf.Unmarshal(cfg)
	if err != nil {
		return err
	}

	if !conf.IsSet(protoHTTP) {
		cfg.HTTP = nil
	} else {
		if cfg.HTTP.Endpoint, err = sanitizeURLPath(cfg.HTTP.Endpoint); err != nil {
			return err
		}
	}

	return nil
}

func sanitizeURLPath(urlPath string) (string, error) {
	u, err := url.Parse(urlPath)
	if err != nil {
		return "", fmt.Errorf("invalid HTTP URL path set for signal: %w", err)
	}

	if !path.IsAbs(u.Path) {
		u.Path = "/" + u.Path
	}
	return u.Path, nil
}
