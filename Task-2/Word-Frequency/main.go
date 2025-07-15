package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func wordFrequencyCount(str string) map[string]int {
	wordCount := make(map[string]int)

	// Convert to lowercase for case-insensitivity
	str = strings.ToLower(str)

	// Remove punctuation by replacing with space
	var cleaned []rune
	for _, r := range str {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') || r == ' ' {
			cleaned = append(cleaned, r)
		} else {
			cleaned = append(cleaned, ' ')
		}
	}

	words := strings.Fields(string(cleaned))
	for _, word := range words {
		wordCount[word]++
	}

	return wordCount
}

func main() {
	var str string
	fmt.Println("Enter a sentence:")

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
