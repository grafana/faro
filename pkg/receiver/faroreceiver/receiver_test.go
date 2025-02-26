// SPDX-License-Identifier: Apache-2.0

package faroreceiver // import "github.com/grafana/faro/pkg/receiver/faroreceiver"

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component/componenttest"
	"go.opentelemetry.io/collector/consumer/consumertest"
	"go.opentelemetry.io/collector/receiver/receivertest"
	"go.uber.org/zap"
)

func TestFaroReceiver_Start(t *testing.T) {
	testcases := []struct {
		name               string
		payload            string
		expectedStatusCode int
		expectedTraces     int
	}{
		{
			name:               "empty",
			payload:            "testdata/empty.json",
			expectedStatusCode: http.StatusBadRequest,
			expectedTraces:     0,
		},
		{
			name:               "minimal-traces-only",
			payload:            "testdata/minimal-traces-only.json",
			expectedStatusCode: http.StatusOK,
			expectedTraces:     1,
		},
	}

	cfg := createDefaultConfig().(*Config)
	logger, err := zap.NewDevelopment()
	require.NoError(t, err)
	defer func() {
		_ = logger.Sync()
	}()

	settings := receivertest.NewNopSettings()
	settings.Logger = logger
	receiver, err := newFaroReceiver(cfg, &settings)
	require.NoError(t, err)

	nextTraces := new(consumertest.TracesSink)
	receiver.RegisterTracesConsumer(nextTraces)
	nextLogs := new(consumertest.LogsSink)
	receiver.RegisterLogsConsumer(nextLogs)

	host := componenttest.NewNopHost()
	err = receiver.Start(context.Background(), host)
	require.NoError(t, err)
	defer func() { require.NoError(t, receiver.Shutdown(context.Background())) }()

	server := httptest.NewServer(http.HandlerFunc(receiver.handleFaroRequest))
	defer server.Close()

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodPost, server.URL+faroPath, nil)
			require.NoError(t, err)

			req.Header.Set("Content-Type", "application/json")

			file, err := os.Open(tc.payload)
			require.NoError(t, err)
			defer file.Close()
			jsonPayload, err := io.ReadAll(file)
			require.NoError(t, err)
			req.Body = io.NopCloser(bytes.NewBuffer(jsonPayload))

			resp, err := http.DefaultClient.Do(req)
			require.NoError(t, err)
			defer resp.Body.Close()

			assert.Equal(t, tc.expectedStatusCode, resp.StatusCode)

			traces := nextTraces.AllTraces()
			assert.Len(t, traces, tc.expectedTraces)
		})
	}
}
