package main

import (
	"fmt"
	"math/rand"
	"strings"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// To Do:  Add ints to the letters
// clean up and test main
func main() {
	urlCodes := make(map[string]string)
	url := "https://gobyexample.com/"
	code, exists := CheckCodeExists(url, urlCodes)
	len := 5
	if !exists {
		GenerateCode(len)
	}
	// Add to urlCodes
	fmt.Print(code)
	fmt.Print(exists)
}

func CheckCodeExists(url string, urlCodes map[string]string) (string, bool) {
	val, ok := urlCodes[url]
	return val, ok
}

// Generating a unique code for a given URL
func GenerateCode(len int) string {
	code := make([]string, len)
	for i := range code {
		code[i] = string(letters[rand.Intn(52)])
	}
	c := strings.Join(code, "")

	return c
}

func AddCode(url string, code string, urlCodes map[string]string) map[string]string {
	urlCodes[url] = code
	return urlCodes
}
