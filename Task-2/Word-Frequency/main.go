package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func wordFrequencyCount(str string) map[string]int {
	wordCount := make(map[string]int)

	words := strings.Fields(str)
	for _, word := range words {
		wordCount[word]++
	}
	return wordCount
}

func main() {
	var str string
	fmt.Println("Enter a sentence:")
	
	// Read the full line including spaces
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		str = scanner.Text()
	}

	counts := wordFrequencyCount(str)
	fmt.Println("Word Frequency Count:")
	for word, count := range counts {
		fmt.Printf("%s: %d\n", word, count)
	}
}
