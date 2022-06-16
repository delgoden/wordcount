package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	phrase := os.Args[1]
	count := 0
	if phrase != "" {
		count = len(strings.Split(phrase, " "))
	}
	fmt.Println(count)
}
