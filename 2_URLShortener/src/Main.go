package main

import (
	"fmt"
	"net/http"
	"urlshortener"
)

func main() {
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
		"/headless-tekla": "https://github.com/TrimbleSolutionsCorporation/HeadlessTeklaStructuresExample/tree/2018",
		"/gophercises":    "https://gophercises.com/",
		"/gowebdev":       "https://www.usegolang.com/",
	}

	mapHandler := urlshortener.MapHandler(pathsToUrls, mux)

	// 	// Build the YAMLHandler using the mapHandler as the
	// 	// fallback
	// 	yaml := `
	// - path: /urlshort
	//   url: https://github.com/gophercises/urlshort
	// - path: /urlshort-final
	//   url: https://github.com/gophercises/urlshort/tree/solution
	// `
	// 	yamlHandler, err := urlshort.YAMLHandler([]byte(yaml), mapHandler)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	fmt.Println("Starting the server on :8080")
	// 	http.ListenAndServe(":8080", yamlHandler)

	fmt.Println("Listening on port :8080")
	http.ListenAndServe(":8080", mapHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
