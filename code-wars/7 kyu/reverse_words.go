package main

import (
	"fmt"
	"strings"
)

func ReverseWords(str string) string {
	words := strings.Split(str, " ")
	reversedWords := make([]string, 0)

	for _, word := range words {
		reversedWord := ""

		for i := len(string(word)) - 1; i >= 0; i-- {
			reversedWord += string(string(word)[i])
		}

		reversedWords = append(reversedWords, reversedWord)
	}

	return strings.Join(reversedWords, " ")
}

func main() {
	fmt.Println(ReverseWords("This is an example!"))
	fmt.Println(ReverseWords("double  spaces"))
}
