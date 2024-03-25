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

func TestAddCode(t *testing.T) {
	tests := []struct {
		testNumber       int
		url              string
		code             string
		urlCodes         map[string]string
		expectedUrlCodes map[string]string
	}{
		{1, "https://gobyexample.com/", "gbecO", map[string]string{}, map[string]string{"https://gobyexample.com/": "gbecO"}},
		{2, "https://courses.calhoun.io/lessons/les_goph_03", "ccill", map[string]string{"https://courses.calhoun.io/lessons/les_goph_04": "ccillg4"}, map[string]string{"https://courses.calhoun.io/lessons/les_goph_04": "ccillg4", "https://courses.calhoun.io/lessons/les_goph_03": "ccill"}},
	}
	for _, test := range tests {
		urlCodesOutput := AddCode(test.url, test.code, test.urlCodes)

		msg := fmt.Sprintf("Expected a match in test %d.", test.testNumber)
		assert.Equal(t, test.expectedUrlCodes, urlCodesOutput, msg)
		assert.Equal(t, test.expectedUrlCodes, urlCodesOutput, msg)
	}
}
