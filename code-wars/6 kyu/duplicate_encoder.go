package main

import (
	"fmt"
	"strings"
)

func CountChar(str string) map[string]int {
	charCount := make(map[string]int)

	for _, char := range str {
		lowerCaseChar := strings.ToLower(string(char))
		charCount[lowerCaseChar] += 1
	}

	return charCount
}

func DuplicateEncode(word string) string {
	encodedString := ""
	charCount := CountChar(word)

	for _, char := range word {
		lowerCaseChar := strings.ToLower(string(char))

		if charCount[lowerCaseChar] >= 2 {
			encodedString += ")"
		} else {
			encodedString += "("
		}
	}

	return encodedString
}

func main() {
	fmt.Println(DuplicateEncode("din"))
	fmt.Println(DuplicateEncode("recede"))
	fmt.Println(DuplicateEncode("Success"))
}
