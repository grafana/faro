// SPDX-License-Identifier: Apache-2.0

package faroexporter // import "github.com/grafana/faro/pkg/exporter/faroexporter"

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component/componenttest"
	"go.opentelemetry.io/collector/config/confighttp"
	"go.opentelemetry.io/collector/exporter/exportertest"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/ptrace"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
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

func TestConsumeTraces(t *testing.T) {
	tests := []struct {
		name           string
		responseStatus int
		wantErr        assert.ErrorAssertionFunc
		numOfRequests  int32
		td             ptrace.Traces
	}{
		{
			name:           "empty",
			responseStatus: http.StatusAccepted,
			wantErr:        assert.NoError,
			numOfRequests:  0,
			td: func() ptrace.Traces {
				traces := ptrace.NewTraces()
				return traces
			}(),
		},
		{
			name:           "two spans with the same resource",
			responseStatus: http.StatusAccepted,
			wantErr:        assert.NoError,
			numOfRequests:  1,
			td: func() ptrace.Traces {
				traces := ptrace.NewTraces()
				resSpans := traces.ResourceSpans().AppendEmpty()
				resSpans.Resource().Attributes().PutStr(string(semconv.ServiceNameKey), "example-value")

				scopeSpans := resSpans.ScopeSpans().AppendEmpty()
				scopeSpans.Spans().AppendEmpty().Attributes().PutStr("example-name", "example-value")
				scopeSpans.Spans().AppendEmpty().Attributes().PutStr("example-name-1", "example-value-1")

				return traces
			}(),
		},
		{
			name:           "two spans with different resources",
			responseStatus: http.StatusAccepted,
			wantErr:        assert.NoError,
			numOfRequests:  2,
			td: func() ptrace.Traces {
				traces := ptrace.NewTraces()
				resSpans1 := traces.ResourceSpans().AppendEmpty()
				resSpans1.Resource().Attributes().PutStr(string(semconv.ServiceNameKey), "example-value")
				scopeSpans := resSpans1.ScopeSpans().AppendEmpty()
				scopeSpans.Spans().AppendEmpty().Attributes().PutStr("example-name", "example-value")

				resSpans2 := traces.ResourceSpans().AppendEmpty()
				resSpans2.Resource().Attributes().PutStr(string(semconv.ServiceNameKey), "example-value-1")
				scopeSpans2 := resSpans2.ScopeSpans().AppendEmpty()
				scopeSpans2.Spans().AppendEmpty().Attributes().PutStr("example-name", "example-value")

				return traces
			}(),
		},
		{
			name:           "two spans with different resources when server returns error",
			responseStatus: http.StatusBadRequest,
			wantErr:        assert.Error,
			numOfRequests:  2,
			td: func() ptrace.Traces {
				traces := ptrace.NewTraces()
				resSpans1 := traces.ResourceSpans().AppendEmpty()
				resSpans1.Resource().Attributes().PutStr(string(semconv.ServiceNameKey), "example-value")
				scopeSpans := resSpans1.ScopeSpans().AppendEmpty()
				scopeSpans.Spans().AppendEmpty().Attributes().PutStr("example-name", "example-value")

				resSpans2 := traces.ResourceSpans().AppendEmpty()
				resSpans2.Resource().Attributes().PutStr(string(semconv.ServiceNameKey), "example-value-1")
				scopeSpans2 := resSpans2.ScopeSpans().AppendEmpty()
				scopeSpans2.Spans().AppendEmpty().Attributes().PutStr("example-name", "example-value")

				return traces
			}(),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var numOfRequests int32
			srv := createBackend("/faro", func(writer http.ResponseWriter, _ *http.Request) {
				numOfRequests++
				writer.WriteHeader(test.responseStatus)
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
			err = expT.ConsumeTraces(context.Background(), test.td)
			test.wantErr(t, err)
			require.Equal(t, test.numOfRequests, numOfRequests)
		})
	}
}

func TestConsumeLogs(t *testing.T) {
	tests := []struct {
		name           string
		responseStatus int
		wantErr        assert.ErrorAssertionFunc
		numOfRequests  int32
		ld             plog.Logs
	}{
		{
			name:           "empty",
			responseStatus: http.StatusAccepted,
			wantErr:        assert.NoError,
			numOfRequests:  0,
			ld: func() plog.Logs {
				logs := plog.NewLogs()
				return logs
			}(),
		},
		{
			name:           "two log records with the same resource",
			responseStatus: http.StatusAccepted,
			wantErr:        assert.NoError,
			numOfRequests:  1,
			ld: func() plog.Logs {
				logs := plog.NewLogs()
				resLogs := logs.ResourceLogs().AppendEmpty()
				resLogs.Resource().Attributes().PutStr(string(semconv.ServiceNameKey), "example-value")

				records := resLogs.ScopeLogs().AppendEmpty().LogRecords()
				records.AppendEmpty().Body().SetStr("kind=event")
				records.AppendEmpty().Body().SetStr("kind=measurement")

				return logs
			}(),
		},
		{
			name:           "two log records with different resources",
			responseStatus: http.StatusAccepted,
			wantErr:        assert.NoError,
			numOfRequests:  2,
			ld: func() plog.Logs {
				logs := plog.NewLogs()
				resLogs1 := logs.ResourceLogs().AppendEmpty()
				resLogs1.Resource().Attributes().PutStr(string(semconv.ServiceNameKey), "example-value")

				records := resLogs1.ScopeLogs().AppendEmpty().LogRecords()
				records.AppendEmpty().Body().SetStr("kind=event")

				resLogs2 := logs.ResourceLogs().AppendEmpty()
				resLogs2.Resource().Attributes().PutStr(string(semconv.ServiceNameKey), "example-value-1")

				records2 := resLogs2.ScopeLogs().AppendEmpty().LogRecords()
				records2.AppendEmpty().Body().SetStr("kind=event")

				return logs
			}(),
		},
		{
			name:           "two log records with different resources when server returns error",
			responseStatus: http.StatusBadRequest,
			wantErr:        assert.Error,
			numOfRequests:  2,
			ld: func() plog.Logs {
				logs := plog.NewLogs()
				resLogs1 := logs.ResourceLogs().AppendEmpty()
				resLogs1.Resource().Attributes().PutStr(string(semconv.ServiceNameKey), "example-value")

				records := resLogs1.ScopeLogs().AppendEmpty().LogRecords()
				records.AppendEmpty().Body().SetStr("kind=event")

				resLogs2 := logs.ResourceLogs().AppendEmpty()
				resLogs2.Resource().Attributes().PutStr(string(semconv.ServiceNameKey), "example-value-1")

				records2 := resLogs2.ScopeLogs().AppendEmpty().LogRecords()
				records2.AppendEmpty().Body().SetStr("kind=event")

				return logs
			}(),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var numOfRequests int32
			srv := createBackend("/faro", func(writer http.ResponseWriter, _ *http.Request) {
				numOfRequests++
				writer.WriteHeader(test.responseStatus)
			})
			defer srv.Close()

			cfg := &Config{
				ClientConfig: confighttp.ClientConfig{
					Endpoint: srv.URL + "/faro",
				},
			}
			expL, err := createLogs(context.Background(), exportertest.NewNopSettings(), cfg)
			require.NoError(t, err)
			err = expL.Start(context.Background(), componenttest.NewNopHost())
			require.NoError(t, err)
			t.Cleanup(func() {
				require.NoError(t, expL.Shutdown(context.Background()))
			})
			err = expL.ConsumeLogs(context.Background(), test.ld)
			test.wantErr(t, err)
			require.Equal(t, test.numOfRequests, numOfRequests)
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
