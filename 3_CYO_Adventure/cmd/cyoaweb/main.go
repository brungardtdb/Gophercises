package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"log"
	"net/http"

	"GOPHERCISES/3_CYO_Adventure/cyoa"
)

func main() {

	absPath, err := filepath.Abs("../../gopher.json")
	check(err)
	portName := flag.Int("port", 8080, "the port to start the server on")
	fileName := flag.String("file", absPath, "the JSON file containing the CYOA story")
	flag.Parse()
	fmt.Printf("Using the story in %s.\n", *fileName)

	f, err := os.Open(*fileName)
	check(err)

	story, err := cyoa.JsonStory(f)
	check(err)

	h := cyoa.NewHandler(story/*, cyoa.WithTemplate(nil)*/) // example of passing in options for handler
	port := fmt.Sprintf(":%v", *portName)
	fmt.Println("Starting the server on port ", port)
	log.Fatal(http.ListenAndServe(port, h))
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
