/*
Create a quiz game using a csv file that asks the user to answer questions
The first version of this quiz will not be timed, but the second version will be
*/

package main

import (
	"encoding/csv"
	"os"
)

func main() {

	file, err := os.Open("quiz.csv")
	check(err)
	defer file.Close()
	parseFile(file)
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
