package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

//  To Do: Questions given invalid answers are considered incorrect.

func main() {
	quiz := import_quiz()
	_, quiz_map_str := compose_questions(quiz)
	quiz_map_int := convert_answers(quiz_map_str)
	successes, failures := perform_quiz(quiz_map_int)
	output_results(successes, failures)
	fmt.Print(successes, failures)
}

// At the end of the quiz the program should output the total number of questions correct and how many questions there were in total
func output_results(successes int, failures int) {
	results := fmt.Sprint("Thank you for taking the quiz! /n Here are your results /n Successes: %i Failures: %i", successes, failures)
	fmt.Print(results)
}

// will then give the quiz to a user keeping track of how many questions they get right and how many they get incorrect.
// Regardless of whether the answer is correct or wrong the next question should be asked immediately afterwards.
func perform_quiz(quiz_map_int map[string]int) (int, int) {
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
			fmt.Print("Nope :/  \n")
			failures += 1
		}
	}
	return successes, failures
}

func convert_answers(quiz_map_str map[string]string) map[string]int {
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

func compose_questions(quiz string) ([]string, map[string]string) {
	entries := strings.Fields(quiz)
	quiz_map_str := make(map[string]string)

	for _, entry := range entries {
		q_and_a := strings.Split(entry, ",")
		q := q_and_a[0]
		a := q_and_a[1]
		quiz_map_str[q] = a
	}
	return entries, quiz_map_str
}

// Create a program that will read in a quiz provided via a CSV file
func import_quiz() string {
	dat, err := os.ReadFile("quiz.csv")
	check(err)
	return string(dat)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
