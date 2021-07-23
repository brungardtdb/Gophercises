package main

import (
	"GOPHERCISES/2_URLShortener/urlshortener"
	"fmt"
	"net/http"
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

	// Build the YAMLHandler using the mapHandler as the
	// fallback
	yaml := `
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
`

	yamlHandler, err := urlshortener.YAMLHandler([]byte(yaml), mapHandler)
	if err != nil {
		panic(err)
	}

	// Build the JSONHandler using the YAMLHandler as the fallback
	json := `
		[{"path":"/go-json","url":"https://golang.org/pkg/encoding/json/"},
		{"path":"/go-json-byexample","url":"https://gobyexample.com/json"},
		{"path":"/hipsteripsum","url":"https://hipsum.co/"},
		{"path":"/json","url":"https://www.json.org/json-en.html"}]`

	jsonHandler, err := urlshortener.JSONHandler([]byte(json), yamlHandler)
	if err != nil {
		panic(err)
	}

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", jsonHandler)

}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
