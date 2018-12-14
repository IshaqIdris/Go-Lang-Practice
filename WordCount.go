package main

import (
	"fmt"
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	fmt.Println(words)
	wordCount := make(map[string]int)
	for i := 0; i < len(words); i++ {
		fmt.Println(words[i])
		wordCount[words[i]]++
	}
	return wordCount
}

func main() {
	wc.Test(WordCount)
}
