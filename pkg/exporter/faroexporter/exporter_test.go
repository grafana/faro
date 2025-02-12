// SPDX-License-Identifier: Apache-2.0

package faroexporter // import "github.com/grafana/faro/pkg/exporter/faroexporter"

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component/componenttest"
	"go.opentelemetry.io/collector/config/confighttp"
	"go.opentelemetry.io/collector/exporter/exportertest"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/ptrace"
)

func TestAcceptedResponsesAndFormats(t *testing.T) {
	tests := []struct {
		name           string
		responseStatus int
		responseBody   string
		err            func(srv *httptest.Server) error
		isPermErr      bool
		headers        map[string]string
	}{
		{
			name:           "202",
			responseStatus: http.StatusAccepted,
			responseBody:   "",
			isPermErr:      true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			srv := createBackend("/faro", func(writer http.ResponseWriter, _ *http.Request) {
				for k, v := range test.headers {
					writer.Header().Add(k, v)
				}
				writer.WriteHeader(test.responseStatus)
				writer.Write([]byte(test.responseBody))
			})
			defer srv.Close()

			cfg := &Config{
				ClientConfig: confighttp.ClientConfig{
					Endpoint: srv.URL + "/faro",
				},
			}

			expT, err := createTraces(context.Background(), exportertest.NewNopSettings(), cfg)
			require.NoError(t, err)

			err = expT.Start(context.Background(), componenttest.NewNopHost())
			require.NoError(t, err)
			t.Cleanup(func() {
				require.NoError(t, expT.Shutdown(context.Background()))
			})

			traces := ptrace.NewTraces()
			err = expT.ConsumeTraces(context.Background(), traces)
			require.NoError(t, err)

			expL, err := createLogs(context.Background(), exportertest.NewNopSettings(), cfg)
			require.NoError(t, err)

			err = expL.Start(context.Background(), componenttest.NewNopHost())
			require.NoError(t, err)
			t.Cleanup(func() {
				require.NoError(t, expL.Shutdown(context.Background()))
			})

			logs := plog.NewLogs()
			err = expL.ConsumeLogs(context.Background(), logs)
			require.NoError(t, err)
		})
	}
}

func createBackend(endpoint string, handler func(writer http.ResponseWriter, request *http.Request)) *httptest.Server {
	mux := http.NewServeMux()
	mux.HandleFunc(endpoint, handler)
	srv := httptest.NewServer(mux)
	fmt.Printf("Server URL: %s\n", srv.URL)
	return srv
}
