package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckCodeExists(t *testing.T) {
	tests := []struct {
		url          string
		urlCodes     map[string]string
		codeOutput   string
		existsOutput bool
	}{
		{"https://gobyexample.com/", map[string]string{}, "", false},
		{"https://courses.calhoun.io/lessons/les_goph_04", map[string]string{"https://courses.calhoun.io/lessons/les_goph_04": "ccillg4"}, "ccillg4", true},
	}
	for _, test := range tests {
		url, exists := CheckCodeExists(test.url, test.urlCodes)
		msg := fmt.Sprintf("%s has not yet been shortened so should return %s and %t rather than %s and %t.", test.url, test.codeOutput, test.existsOutput, url, exists)
		assert.Equal(t, url, test.codeOutput, msg)
		assert.Equal(t, exists, test.existsOutput, msg)
	}
}

func TestGenerateCode(t *testing.T) {
	tests := []struct {
		len int
	}{
		{5},
		{6},
	}
	for _, test := range tests {
		code := GenerateCode(test.len)

		msg := fmt.Sprintf("%s is not %d characters .", code, test.len)
		assert.Equal(t, test.len, len(code), msg)
		assert.Equal(t, test.len, len(code), msg)
	}
}
