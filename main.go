package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	phrase := os.Args[1]
	wordSls := strings.Split(phrase, " ")
	fmt.Println(len(wordSls))
}
