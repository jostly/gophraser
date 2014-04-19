package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jostly/gophraser/dict"
	"github.com/jostly/gophraser/response"
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
	phrase := response.Phrase{adj, animals.OneStartingWith(adj[:1])}

	fmt.Fprintln(res, response.BuildResponse(req, phrase, res))
}

func letterPhraser(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	letter := vars["letter"]
	phrase := response.Phrase{adjectives.OneStartingWith(letter), animals.OneStartingWith(letter)}
	fmt.Fprintln(res, response.BuildResponse(req, phrase, res))
}

func randomPhraser(res http.ResponseWriter, req *http.Request) {
	phrase := response.Phrase{adjectives.OneRandom(), animals.OneRandom()}
	fmt.Fprintln(res, response.BuildResponse(req, phrase, res))
}
