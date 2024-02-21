package main

import "fmt"

func main() {
	s := "Hello World"
	chopped_string := StringChop(s)
	fmt.Print(chopped_string)
}

func StringChop(s string) string {
	// a function that removes the first and last characters from a string.
	length := len(s)
	r := []rune(s)
	end_index := length - 1
	return string(append(r[1:end_index]))
}

func CheckLength(numbers []int) bool {
	// Return true if the marathon is 25 miles long, otherwise, return false.
	sum := 0

	for _, num := range numbers {
		if num < 0 {
			sum -= num // This way, we don't convert to float and back, which is more efficient.
		} else {
			sum += num
		}
	}

	return sum == 25
}
