package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/rtravitz/autocomplete/trie"
)

func main() {
	dat, err := ioutil.ReadFile("./words")
	if err != nil {
		log.Fatal(err)
	}
	words := strings.Split(string(dat), "\n")
	autocomplete := trie.New()
	autocomplete.Build(words)
	suggestions := autocomplete.Suggest("dog")
	fmt.Println(suggestions)
}
