/*
Create a quiz game using a csv file that asks the user to answer questions
The first version of this quiz will not be timed, but the second version will be
The program will keep track of all right and wrong answers from the user
*/

package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	// "time"
)

func main() {

	file, err := os.Open("quiz.csv")
	check(err)
	defer file.Close()
	fileData := parseFile(file)
	quizUser(fileData)
}

func askQuestion(question []string) chan bool {

	answer := make(chan bool)
	go func() {
		// time.Sleep(10 * time.Second)
		fmt.Println("Please answer the following question:")
		fmt.Println(question[0])
		// Convert answer from csv to int
		intAnswer, err := strconv.Atoi(question[1])
		check(err)
		// Prompt user for answer and convert to int
		var intGuess int
		_, err = fmt.Scan(&intGuess)
		check(err)
		answer <- intGuess == intAnswer
	}()

	return answer
}

func quizUser(questions [][]string) {

	correctGuesses := 0
	incorrectGuesses := 0

	// Iterate through questions from csv file, quizzing user
	for i := 0; i < len(questions); i++ {

		question := questions[i]
		answerIsCorrect := askQuestion(question)
		// Total up correct and incorrect guesses
		if <-answerIsCorrect {
			correctGuesses++
		} else {
			incorrectGuesses++
		}
	}

	fmt.Println("You guessed " + strconv.Itoa(correctGuesses) + " correctly!")
	fmt.Println("You guessed " + strconv.Itoa(incorrectGuesses) + " incorrectly!")
}

func parseFile(file *os.File) [][]string {

	reader := csv.NewReader(file)
	fileData, err := reader.ReadAll()
	check(err)
	return fileData
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
