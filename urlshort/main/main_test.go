package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGreet(t *testing.T) {
	tests := []struct {
		expectedOutput string
	}{
		{"Hello World"},
	}
	for _, test := range tests {
		output := Greet()
		msg := fmt.Sprintf("%s should match %s.", output, test.expectedOutput)
		assert.Equal(t, output, test.expectedOutput, msg)
	}
}

func TestCheckCodeExists(t *testing.T) {
	tests := []struct {
		url          string
		codeOutput   string
		existsOutput bool
	}{
		{"https://gobyexample.com/", "", false},
	}
	for _, test := range tests {
		url, exists := CheckCodeExists(test.url)
		msg := fmt.Sprintf("%s has not yet been shortened so should return %s and %t rather than %s and %t.", test.url, test.codeOutput, test.existsOutput, url, exists)
		assert.Equal(t, url, test.codeOutput, msg)
		assert.Equal(t, exists, test.existsOutput, msg)
	}
}
