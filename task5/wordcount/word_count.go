package main

import (
	"fmt"
	"strings"
)

func wordFrequency(text string) map[string]int {
	wordCount := make(map[string]int)

	words := strings.Fields(strings.ToLower(text))

	for _, word := range words {
		wordCount[word]++
	}

	return wordCount
}

func main() {
	text := "Go is easy to learn. Go is fast. Go is fun!"

	counts := wordFrequency(text)

	fmt.Println("Частота слов в тексте:")
	for word, count := range counts {
		fmt.Printf("%s: %d\n", word, count)
	}
}