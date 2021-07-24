package handler

import(
	"fmt"
	"net/http"
	"log"
)

func Run() {

	http.HandleFunc("/intro", intro)
	http.HandleFunc("/newyork", newYork)
	http.HandleFunc("/debate", debate)
	http.HandleFunc("/sean-kelly", seanKelly)
	http.HandleFunc("/mark-bates", markBates)
	http.HandleFunc("/denver", denver)
	http.HandleFunc("/home", home)
	http.HandleFunc("/", notFound)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func intro(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Intro")
}


func newYork(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "New York")
}

func debate(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Debate")
}

func seanKelly(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Sean Kelly")
}

func markBates(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Mark Bates")
}

func denver(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Denver")
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home")
}

func notFound(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Are you lost?")
}

