// SPDX-License-Identifier: Apache-2.0

package faroexporter // import "github.com/grafana/faro/exporter"

import (
	"errors"

	"go.uber.org/multierr"
)

// FaroExporterConfig contains the configuration options for the faro exporter
type FaroExporterConfig struct {
	Endpoint string `mapstructure:"endpoint"`
}

// Config contains the main configuration options for the faro exporter
type Config struct {
	FaroExporter FaroExporterConfig `mapstructure:"faroexporter"`
}

func (c *Config) Validate() error {
	var errs error
	if c.FaroExporter.Endpoint == "" {
		errs = multierr.Append(errs, errors.New("endpoint is required"))
	}
	return errs
}
