package main

import (
	// "flag"
	// "fmt"
	// "os"
	// "path/filepath"

	// "GOPHERCISES/3_CYO_Adventure/cyoa"
	"GOPHERCISES/3_CYO_Adventure/cmd/cyoaweb/handler"
)

func main() {

	// absPath, err := filepath.Abs("../../gopher.json")
	// check(err)
	// fileName := flag.String("file", absPath, "the JSON file containing the CYOA story")
	// flag.Parse()
	// fmt.Printf("Using the story in %s.\n", *fileName)

	// f, err := os.Open(*fileName)
	// check(err)

	// story, err := cyoa.JsonStory(f)
	// check(err)

	// fmt.Printf("%v\n", story)

	handler.Run()
}

// func check(err error) {
// 	if err != nil {
// 		panic(err)
// 	}
// }
