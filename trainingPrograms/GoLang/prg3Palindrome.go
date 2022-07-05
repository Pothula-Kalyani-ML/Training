package main

import (
	"fmt"
)

func palindrome(inputString string) bool {
	var firstIndex = 0
	var lastIndex = len([]rune(inputString)) - 1
	if lastIndex <= 0 {
		fmt.Println("not palindrome")
		return false
	}
	for firstIndex < lastIndex {
		if inputString[firstIndex] != inputString[lastIndex] {
			fmt.Println("not palindrome")
			return false
		}
		firstIndex++
		lastIndex--
	}
	return true

}

func main() {
	if palindrome(" ") {
		fmt.Println("palindrome")
	}

}
