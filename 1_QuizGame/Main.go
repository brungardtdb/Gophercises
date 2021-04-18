/*
Create a quiz game using a csv file that asks the user to answer questions
The first version of this quiz will not be timed, but the second version will be
The program will keep track of all right and wrong answers from the user
*/

package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type problem struct {
	question string
	answer   string
}

func main() {

	csvFileName := flag.String("csv", "quiz.csv", "A csv file with a question,answer format")
	flag.Parse()
	_ = csvFileName
	file, err := os.Open(*csvFileName)
	check(err)
	defer file.Close()
	fileData := parseFile(file)
	quizUser(fileData)
}

func quizUser(problems []problem) {

	correctGuesses := 0
	incorrectGuesses := 0
	timeLimit := flag.Int("limit", 30, "The time limit for the quiz in seconds.")
	flag.Parse()
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second) // quiz will end when time limit is up

	// Iterate through questions from csv file, quizzing user
	for i := 0; i < len(problems); i++ {

		userAnswer := make(chan bool, 1)
		newProblem := problems[i]
		fmt.Println(i+1, ". ", newProblem.question)

		go func() {
			// Convert answer from string to int
			intAnswer, err := strconv.Atoi(newProblem.answer)
			check(err)
			// Prompt user for answer and convert to int
			var intGuess int
			_, err = fmt.Scan(&intGuess)
			check(err)
			userAnswer <- intGuess == intAnswer
		}()

		select {
		case <-timer.C: // if we hear from timer, quiz is over
			fmt.Println("\nQuiz time limit exceeded!")
			resultsMessage(correctGuesses, incorrectGuesses, len(problems))
			return
		case problemResult := <-userAnswer: // record correct and incorrect answers
			if problemResult {
				correctGuesses++
			} else {
				incorrectGuesses++
			}
		}
	}
	// print final results
	resultsMessage(correctGuesses, incorrectGuesses, len(problems))
}

func resultsMessage(correctGuesses int, incorrectGuesses int, numQuestions int) {
	fmt.Println("You guessed " + strconv.Itoa(correctGuesses) + " correctly!")
	fmt.Println("You guessed " + strconv.Itoa(incorrectGuesses) + " incorrectly!")
	fmt.Println("Your score is " + strconv.Itoa(correctGuesses) + "/" + strconv.Itoa(numQuestions))
}

func parseFile(file *os.File) []problem {

	reader := csv.NewReader(file)
	fileData, err := reader.ReadAll()
	check(err)

	// Extract problems from csv
	problems := make([]problem, len(fileData))
	for i, line := range fileData {
		problems[i] = problem{
			question: strings.TrimSpace(line[0]),
			answer:   strings.TrimSpace(line[1]),
		}
	}
	return problems
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
