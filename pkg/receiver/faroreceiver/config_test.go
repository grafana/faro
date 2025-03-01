// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package faroreceiver

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/configauth"
	"go.opentelemetry.io/collector/config/confighttp"
	"go.opentelemetry.io/collector/config/configtls"
	"go.opentelemetry.io/collector/confmap/confmaptest"
)

func TestUnmarshalDefaultConfig(t *testing.T) {
	cm, err := confmaptest.LoadConf(filepath.Join("testdata", "default.yaml"))
	require.NoError(t, err)
	factory := NewFactory()
	cfg := factory.CreateDefaultConfig()
	require.NoError(t, cm.Unmarshal(&cfg))
	assert.Equal(t, factory.CreateDefaultConfig(), cfg)
}

func TestUnmarshalConfig(t *testing.T) {
	cm, err := confmaptest.LoadConf(filepath.Join("testdata", "config.yaml"))
	require.NoError(t, err)
	factory := NewFactory()
	cfg := factory.CreateDefaultConfig()
	require.NoError(t, cm.Unmarshal(&cfg))
	assert.Equal(t,
		&Config{
			ServerConfig: &confighttp.ServerConfig{
				Auth: &confighttp.AuthConfig{
					Authentication: configauth.Authentication{
						AuthenticatorID: component.MustNewID("test"),
					},
				},
				Endpoint: "localhost:8080",
				TLSSetting: &configtls.ServerConfig{
					Config: configtls.Config{
						CertFile: "test.crt",
						KeyFile:  "test.key",
					},
				},
				CORS: &confighttp.CORSConfig{
					AllowedOrigins: []string{"https://*.test.com", "https://test.com"},
					MaxAge:         7200,
				},
			},
		}, cfg)
}
