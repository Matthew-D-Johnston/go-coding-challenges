package main

import (
	"fmt"
	"strings"
)

func SpinWords(str string) string {
	words := strings.Split(str, " ")
	spunWords := make([]string, 0)

	for _, word := range words {
		if len(word) >= 5 {
			spunWord := ""

			for i := len(word) - 1; i >= 0; i-- {
				spunWord += string(word[i])
			}

			spunWords = append(spunWords, spunWord)
		} else {
			spunWords = append(spunWords, word)
		}
	}

	return strings.Join(spunWords, " ")
}

func main() {
	fmt.Println(SpinWords("Hey fellow warriors"))
	fmt.Println(SpinWords("This is a test"))
	fmt.Println(SpinWords("This is another test"))
}
