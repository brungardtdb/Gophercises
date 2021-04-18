/*
Create a quiz game using a csv file that asks the user to answer questions
The first version of this quiz will not be timed, but the second version will be
*/

package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

func main() {

	file, err := os.Open("quiz.csv")
	check(err)
	defer file.Close()
	fmt.Println(file)
	fileData := parseFile(file)
	quizUser(fileData)
}

func askQuestion(question []string) chan bool {

	answer := make(chan bool)
	go func() {
		time.Sleep(10 * time.Second)
		answer <- true
	}()

	return answer
}

func quizUser(questions [][]string) {

	for i := 0; i < len(questions); i++ {
		question := questions[i]
		fmt.Println(question[0])
		fmt.Println(question[1])
		staticBool := askQuestion(question)
		fmt.Println(<-staticBool)
	}
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
