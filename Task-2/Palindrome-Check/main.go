package main

import (
	"fmt"
	"strings"
)

func palindrome_check(str string) string {
	var cleaned []rune
	for _, r := range str {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') {
			cleaned = append(cleaned, rune(strings.ToLower(string(r))[0]))
		}
	}
	left := 0
	right := len(cleaned) - 1
	for left < right {
		if cleaned[left] != cleaned[right] {
			return "Not a palindrome"
		}
		left++
		right--
	}
	return "Palindrome"
}

func main() {
	var str string
	fmt.Scanln(&str)
	result := palindrome_check(str)
	fmt.Println(result)

}
