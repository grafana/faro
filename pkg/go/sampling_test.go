package faro

import (
	"encoding/json"
	"testing"
)

// A payload from an SDK that does not sample by rule has no meta.sampling. It must
// decode with Sampling == nil, not with a zero-valued struct, so receivers can tell
// "no verdict" from "a verdict of zeros".
func TestMetaSamplingAbsentIsNil(t *testing.T) {
	var p Payload
	if err := json.Unmarshal([]byte(`{"meta":{"session":{"id":"s0"}}}`), &p); err != nil {
		t.Fatal(err)
	}
	if p.Meta.Sampling != nil {
		t.Fatalf("expected nil Sampling, got %+v", *p.Meta.Sampling)
	}
	out, err := json.Marshal(p.Meta)
	if err != nil {
		t.Fatal(err)
	}
	var m map[string]json.RawMessage
	if err := json.Unmarshal(out, &m); err != nil {
		t.Fatal(err)
	}
	if _, ok := m["sampling"]; ok {
		t.Fatalf("nil Sampling must not be serialised, got %s", out)
	}
}

// The design doc's example verdict, a session kept by the errors trigger.
func TestMetaSamplingDecodes(t *testing.T) {
	body := `{"meta":{"session":{"id":"s0"},"sampling":{
		"schema":1,"sessionId":"s0","inBaseSample":false,"keptBy":"trigger","keptByRule":"errors",
		"baseRate":0.05,"replayRate":null,"triggerRate":null,"replayKept":true,
		"rulesHash":"19aec605c80f4576","keyU32":1262754589}}}`
	var p Payload
	if err := json.Unmarshal([]byte(body), &p); err != nil {
		t.Fatal(err)
	}
	s := p.Meta.Sampling
	if s == nil {
		t.Fatal("expected Sampling, got nil")
	}
	if s.Schema == nil || *s.Schema != 1 {
		t.Errorf("schema: %v", s.Schema)
	}
	if s.SessionID == nil || *s.SessionID != "s0" {
		t.Errorf("sessionId: %v", s.SessionID)
	}
	if s.InBaseSample == nil || *s.InBaseSample {
		t.Errorf("inBaseSample: %v", s.InBaseSample)
	}
	if s.KeptBy == nil || *s.KeptBy != SamplingKeptByTrigger {
		t.Errorf("keptBy: %v", s.KeptBy)
	}
	if s.KeptByRule == nil || *s.KeptByRule != "errors" {
		t.Errorf("keptByRule: %v", s.KeptByRule)
	}
	if s.BaseRate == nil || *s.BaseRate != 0.05 {
		t.Errorf("baseRate: %v", s.BaseRate)
	}
	if s.ReplayRate != nil || s.TriggerRate != nil {
		t.Errorf("null rates must decode as nil: replayRate=%v triggerRate=%v", s.ReplayRate, s.TriggerRate)
	}
	if s.ReplayKept == nil || !*s.ReplayKept {
		t.Errorf("replayKept: %v", s.ReplayKept)
	}
	if s.RulesHash == nil || *s.RulesHash != "19aec605c80f4576" {
		t.Errorf("rulesHash: %v", s.RulesHash)
	}
	if s.KeyU32 == nil || *s.KeyU32 != 1262754589 {
		t.Errorf("keyU32: %v", s.KeyU32)
	}
}

// A partial verdict still decodes; deciding that it is invalid is the receiver's job.
func TestMetaSamplingPartialDecodes(t *testing.T) {
	var p Payload
	if err := json.Unmarshal([]byte(`{"meta":{"sampling":{"schema":1}}}`), &p); err != nil {
		t.Fatal(err)
	}
	s := p.Meta.Sampling
	if s == nil || s.Schema == nil || *s.Schema != 1 {
		t.Fatalf("expected schema 1, got %+v", s)
	}
	if s.SessionID != nil || s.KeyU32 != nil {
		t.Errorf("missing fields must be nil: %+v", *s)
	}
}
