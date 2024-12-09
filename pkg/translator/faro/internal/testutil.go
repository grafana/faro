package internal

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	faro "github.com/grafana/faro/pkg/go"
)

func PayloadFromFile(t *testing.T, filename string) *faro.Payload {
	t.Helper()

	f, err := os.Open(filepath.Join("testdata", filename))
	if err != nil {
		t.Fatal(err)
	}

	defer f.Close()

	var p faro.Payload
	if err := json.NewDecoder(f).Decode(&p); err != nil {
		t.Fatal(err)
	}

	return &p
}
