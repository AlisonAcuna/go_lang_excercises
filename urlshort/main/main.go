package main

import (
	"fmt"
)

var UrlCodes map[string]string

func main() {
	greeting := Greet()
	fmt.Print(greeting)
}

func Greet() string {
	return "Hello World"
}

func CheckCodeExists(url string) (string, bool) {
	val, ok := UrlCodes[url]
	return val, ok
}

// Generating a unique code for a given URL
// func GenerateCode(url string) string {

// }
