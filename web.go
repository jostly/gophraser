package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jostly/gophraser/dict"
)

var (
	animals    = dict.NewDictionary("words/animals.txt")
	adjectives = dict.NewDictionary("words/adjectives.txt")
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/random", randomPhraser)
	r.HandleFunc("/{letter:[a-wy-z]}", letterPhraser)
	r.HandleFunc("/", alliterativePhraser)
	http.Handle("/", r)

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "5000"
	}
	log.Println("listening on port " + port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
}

func alliterativePhraser(res http.ResponseWriter, req *http.Request) {
	adj := adjectives.OneRandom()
	fmt.Fprintln(res, adj+" "+animals.OneStartingWith(adj[:1]))
}

func letterPhraser(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	letter := vars["letter"]
	fmt.Fprintln(res, adjectives.OneStartingWith(letter)+" "+animals.OneStartingWith(letter))
}

func randomPhraser(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, adjectives.OneRandom()+" "+animals.OneRandom())
}
