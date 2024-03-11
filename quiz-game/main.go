package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// To Do

// Adapt your program from part 1 to add a timer. The default time limit should be 30 seconds, but should also be customizable via a flag.

// Your quiz should stop as soon as the time limit has exceeded. That is, you shouldn’t wait for the user to answer one final questions but should ideally stop the quiz entirely even if you are currently waiting on an answer from the end user.

// Users should be asked to press enter (or some other key) before the timer starts, and then the questions should be printed out to the screen one at a time until the user provides an answer. Regardless of whether the answer is correct or wrong the next question should be asked.

// At the end of the quiz the program should still output the total number of questions correct and how many questions there were in total. Questions given invalid answers or unanswered are considered incorrect.

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
	timeout := time.After(30 * time.Second)
	quizChannel := make(chan bool)

	go func() {
		for key, val := range quiz_map_int {
			prompt := fmt.Sprintf("Answer the following question: %s  ", key)
			fmt.Print(prompt)
			answer := RandomAnswer()
			s, f := EvaluateAnswer(answer, val)
			successes += s
			failures += f
		}
		quizChannel <- true
	}()

	select {
	case <-quizChannel:
		// Completed within timeout.
	case <-timeout:
		// Timeout occurred.
		fmt.Print("Timeout!")
		return successes, failures
	}

	return successes, failures
}

func EvaluateAnswer(answer int, val int) (int, int) {
	if answer == val {
		fmt.Print("Sucess! \n")
		return 1, 0
	} else {
		fmt.Print("Nope :/  \n")
		return 0, 1
	}
}

func RandomAnswer() int {
	time.Sleep(8 * time.Second)
	a := rand.Int() % 20
	fmt.Print(a)
	return a
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
