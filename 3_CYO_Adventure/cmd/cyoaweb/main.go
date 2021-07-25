package main

import (
	"flag"
	"path/filepath"

	"GOPHERCISES/3_CYO_Adventure/cmd/cyoaweb/handler"
)

func main() {

	absPath, err := filepath.Abs("../../gopher.json")
	check(err)
	portName := flag.Int("port", 8080, "the port to start the server on")
	fileName := flag.String("file", absPath, "the JSON file containing the CYOA story")
	flag.Parse()
	if err = handler.Run(*portName, *fileName); err != nil {
		panic(err)
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}


