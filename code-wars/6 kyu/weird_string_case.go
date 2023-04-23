package main

import (
	"fmt"
	"strings"
)

func toWeirdCase(str string) string {
	words := strings.Split(str, " ")
	weirdCaseWords := make([]string, 0)

	for _, word := range words {
		weirdCaseWord := ""

		for idx, char := range word {
			if idx%2 == 0 {
				weirdCaseWord += strings.ToUpper(string(char))
			} else {
				weirdCaseWord += strings.ToLower(string(char))
			}
		}

		weirdCaseWords = append(weirdCaseWords, weirdCaseWord)
	}

	return strings.Join(weirdCaseWords, " ")
}

func main() {
	fmt.Println(toWeirdCase("String"))
	fmt.Println(toWeirdCase("Weird string case"))
}
