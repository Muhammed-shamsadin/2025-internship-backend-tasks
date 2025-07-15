package main

import (
	"fmt"
)

func palindrome_check(str string) string {
	left := 0
	right := len(str) - 1

	for left < right {
		if str[left] != str[right] {
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

