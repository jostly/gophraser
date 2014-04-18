package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
)

var (
	animals    = readFile("words/animals.txt")
	adjectives = readFile("words/adjectives.txt")
)

func main() {
	rand.Seed(time.Now().UnixNano())
	http.HandleFunc("/", phraser)
	log.Println("listening...")

	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}
}

func phraser(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, adjectives.OneOf()+" "+animals.OneOf())
}

type Dictionary []string

func (d Dictionary) Contains(s string) bool {
	for _, value := range d {
		if value == s {
			return true
		}
	}
	return false
}

func (d Dictionary) OneOf() string {
	return d[rand.Intn(len(d))]
}

func readFile(filename string) Dictionary {
	result := Dictionary{}

	file, error := os.Open(filename)
	if error != nil {
		log.Fatal(error)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.ToLower(scanner.Text())
		if len(line) > 1 && !result.Contains(line) {
			result = append(result, line)
		}
	}
	return result
}
