package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"GOPHERCISES/3_CYO_Adventure/cyoa"
)

func main() {

	absPath, err := filepath.Abs("../../gopher.json")
	check(err)
	fileName := flag.String("file", absPath, "the JSON file containing the CYOA story")
	flag.Parse()
	fmt.Printf("Using the story in %s.\n", *fileName)

	f, err := os.Open(*fileName)
	check(err)

	d := json.NewDecoder(f)
	var story cyoa.Story

	err = d.Decode(&story)
	check(err)

	fmt.Println(story)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
