package dict

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

func Init() {
	rand.Seed(time.Now().UnixNano())
}

type Dictionary []string

func NewDictionary(filename string) Dictionary {
	return readFile(filename)
}

func (d Dictionary) Contains(s string) bool {
	for _, value := range d {
		if value == s {
			return true
		}
	}
	return false
}

func (d Dictionary) OneRandom() string {
	return d[rand.Intn(len(d))]
}

func (d Dictionary) OneStartingWith(s string) string {
	return d.FilterPrefix(s).OneRandom()
}

func (d Dictionary) FilterPrefix(prefix string) Dictionary {
	result := Dictionary{}
	for _, value := range d {
		if strings.HasPrefix(value, prefix) {
			result = append(result, value)
		}
	}
	return result
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
