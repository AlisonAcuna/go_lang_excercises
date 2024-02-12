package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	quiz := ImportQuiz()
	quiz_map_str := ComposeQuestions(quiz)
	quiz_map_int := ConvertAnswers(quiz_map_str)
	successes, failures := PerformQuiz(quiz_map_int)
	OutputResults(successes, failures)
	fmt.Print(successes, failures)
}

// At the end of the quiz the program should output the total number of questions correct and how many questions there were in total
func OutputResults(successes int, failures int) {
	fmt.Printf("Thank you for taking the quiz! /n Here are your results /n Successes: %d Failures: %d", successes, failures)
}

// will then give the quiz to a user keeping track of how many questions they get right and how many they get incorrect.
// Regardless of whether the answer is correct or wrong the next question should be asked immediately afterwards.
func PerformQuiz(quiz_map_int map[string]int) (int, int) {
	successes := 0
	failures := 0
	for key, val := range quiz_map_int {
		var i int
		prompt := fmt.Sprintf("Answer the following question: %s  ", key)
		fmt.Print(prompt)
		fmt.Scan(&i)
		if i == val {
			fmt.Print("Sucess! \n")
			successes += 1
		} else {
			fmt.Println("Nope :/")
			failures += 1
		}
	}
	return successes, failures
}

func ConvertAnswers(quiz_map_str map[string]string) map[string]int {
	quiz_map_int := make(map[string]int)
	for key, val := range quiz_map_str {
		i, err := strconv.Atoi(val)
		if err != nil {
			panic(err)
		}
		quiz_map_int[key] = i
	}
	return quiz_map_int
}

func ComposeQuestions(quiz string) map[string]string {
	entries := strings.Fields(quiz)
	quiz_map_str := make(map[string]string)

	for _, entry := range entries {
		q_and_a := strings.Split(entry, ",")
		q := q_and_a[0]
		a := q_and_a[1]
		quiz_map_str[q] = a
	}
	return quiz_map_str
}

// Create a program that will read in a quiz provided via a CSV file
func ImportQuiz() string {
	dat, err := os.ReadFile("quiz.csv")
	check(err)
	return string(dat)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
