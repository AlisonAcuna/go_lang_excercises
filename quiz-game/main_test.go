package main

import (
	"bytes"
	"os"
	"testing"
)

func TestOutputResults(t *testing.T) {
	tests := []struct {
		successes  int
		failures   int
		wantOutput string
	}{
		{5, 5, "Thank you for taking the quiz! /n Here are your results /n Successes: 5 Failures: 5"},
	}

	for _, test := range tests {
		oldStdout := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w
		OutputResults(test.successes, test.failures)

		// Close the write-end of the pipe to read the output
		w.Close()
		var buf bytes.Buffer
		buf.ReadFrom(r)
		os.Stdout = oldStdout

		// Get the captured output as a string
		gotOutput := buf.String()

		// Check if the captured output matches the expected output
		if gotOutput != test.wantOutput {
			t.Errorf("OutputResults(%d, %d) = %q, want %q", test.successes, test.failures, gotOutput, test.wantOutput)
		}
	}
}
