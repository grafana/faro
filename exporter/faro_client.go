// SPDX-License-Identifier: Apache-2.0

package faroexporter // import "github.com/grafana/faro/exporter"

import (
	faro "github.com/grafana/faro/pkg/go"
)

type FaroClient interface {
	SendPayload(payload faro.Payload) error
	SendMeasurements(m []faro.Measurement) error
	SendLogs(l []faro.Log) error
	SendTraces(t faro.Traces) error
}

func newFaroClient() *faroClient {
	return &faroClient{}
}

type faroClient struct{}

func (fc *faroClient) SendPayload(payload faro.Payload) error {
	return nil
}

func (fc *faroClient) SendMeasurement(ms []faro.Measurement) error {
	p := faro.Payload{
		Measurements: ms,
	}
	return fc.SendPayload(p)
}

func (fc *faroClient) SendLog(ls []faro.Log) error {
	p := faro.Payload{
		Logs: ls,
	}
	return fc.SendPayload(p)
}

func (fc *faroClient) SendTraces(t faro.Traces) error {
	p := faro.Payload{
		Traces: &t,
	}
	return fc.SendPayload(p)
}
