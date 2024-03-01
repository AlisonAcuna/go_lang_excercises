package main

import "fmt"

func main() {
	fmt.Print("Hello World")
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

// Create a function, that will for a given a, b, c, do the following:

// Add a to itself b times.
// Check if the result is divisible by c.

func MultiDivision(a int, b int, c int) bool {
	for i := 1; i < 10; i++ {
		a = a + a
	}
	return a%c == 0
}

// Create a function that takes two numbers as arguments (num, length) and returns an array of multiples of num until the array length reaches length.
func CreateMultiples(num int, length int) []int {
	results := []int{num}
	sum := num
	for i := 1; i < length; i++ {
		sum = sum + num
		results = append(results, sum)
	}
	return results
}
