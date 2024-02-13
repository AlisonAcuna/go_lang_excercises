package main

import (
	"bytes"
	"os"
	"testing"
)

func TestOutputResults(t *testing.T) {
	tests := []struct {
		successes      int
		failures       int
		expectedOutput string
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
		actualOutput := buf.String()

		// Check if the captured output matches the expected output
		if actualOutput != test.expectedOutput {
			t.Errorf("OutputResults(%d, %d) = %q, want %q", test.successes, test.failures, actualOutput, test.expectedOutput)
		}
	}
}

func TestConvertAnswers(t *testing.T) {
	tests := []struct {
		input          map[string]string
		expectedOutput map[string]int
	}{
		{map[string]string{"8+6": "14", "3+1": "4"}, map[string]int{"8+6": 14, "3+1": 4}},
	}

	for _, test := range tests {
		actualOutput := ConvertAnswers(test.input)

		for key, _ := range actualOutput {
			if actualOutput[key] != test.expectedOutput[key] {
				t.Errorf("ComposeQuestions(%s) = %q, want %q", test.input, actualOutput, test.expectedOutput)
			}
		}
	}
}

func TestComposeQuestions(t *testing.T) {
	tests := []struct {
		quiz           string
		expectedOutput map[string]string
	}{
		{"8+6,14 3+1,4", map[string]string{"8+6": "14", "3+1": "4"}},
	}

	for _, test := range tests {
		actualOutput := ComposeQuestions(test.quiz)
		if len(actualOutput) != len(test.expectedOutput) {
			t.Errorf("ComposeQuestions(%s) = %q, want %q", test.quiz, actualOutput, test.expectedOutput)
		}
		for key, _ := range actualOutput {
			if actualOutput[key] != test.expectedOutput[key] {
				t.Errorf("ComposeQuestions(%s) = %q, want %q", test.quiz, actualOutput, test.expectedOutput)
			}
		}
	}
}
