package faroexporter // import "github.com/grafana/faro/pkg/exporter/faro"

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component/componenttest"
	"go.opentelemetry.io/collector/exporter/exportertest"
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
				FaroEndpoint: srv.URL + "/faro",
			}

			exp, err := createTraces(context.Background(), exportertest.NewNopSettings(), cfg)
			require.NoError(t, err)

			err = exp.Start(context.Background(), componenttest.NewNopHost())
			require.NoError(t, err)
			t.Cleanup(func() {
				require.NoError(t, exp.Shutdown(context.Background()))
			})

			traces := ptrace.NewTraces()
			err = exp.ConsumeTraces(context.Background(), traces)
			require.NoError(t, err)
		})
	}
}

func createBackend(endpoint string, handler func(writer http.ResponseWriter, request *http.Request)) *httptest.Server {
	mux := http.NewServeMux()
	mux.HandleFunc(endpoint, handler)
	srv := httptest.NewServer(mux)
	return srv
}
