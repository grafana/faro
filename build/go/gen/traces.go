package faro

import "go.opentelemetry.io/collector/pdata/ptrace"

// UnmarshalJSON unmarshals Traces model.
func (t *Traces) UnmarshalJSON(b []byte) error {
	unmarshaler := &ptrace.JSONUnmarshaler{}
	td, err := unmarshaler.UnmarshalTraces(b)
	if err != nil {
		return err
	}
	*t = Traces{td}
	return nil
}

// MarshalJSON marshals Traces model to json.
func (t Traces) MarshalJSON() ([]byte, error) {
	marshaler := &ptrace.JSONMarshaler{}
	return marshaler.MarshalTraces(t.Traces)
}
