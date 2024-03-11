package main_test

import (
	"bytes"
	"os"
	"testing"
)

func TestMainOutput(t *testing.T) {
	// Instead of replacing os.Stdout globally, which is not concurrency-safe,
	// we define a new function local to the test that captures stdout.
	mainOutput := func() string {
		old := os.Stdout // keep backup of the real stdout
		r, w, _ := os.Pipe()
		os.Stdout = w

		main()

		w.Close()
		os.Stdout = old // restoring the real stdout
		var buf bytes.Buffer
		if _, err := buf.ReadFrom(r); err != nil {
			t.Fatal(err)
		}
		return buf.String()
	}

	tests := []struct {
		wantOutput string
	}{
		{"hello world\n"}, // We need to include the newline character since fmt.Println adds one.
	}

	for _, test := range tests {
		if gotOutput := mainOutput(); gotOutput != test.wantOutput {
			t.Errorf("Expected output %q, but got %q", test.wantOutput, gotOutput)
		}
	}
}
