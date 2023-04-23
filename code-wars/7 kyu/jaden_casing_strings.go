package main

import (
	"fmt"
	"regexp"
	"strings"
)

func ToJadenCase(str string) string {
	words := strings.Split(str, " ")
	jadenWords := make([]string, 0)

	for _, word := range words {
		firstLetter := strings.ToUpper(string(word[0]))
		re := regexp.MustCompile(`\A.`)
		word = re.ReplaceAllLiteralString(word, firstLetter)
		jadenWords = append(jadenWords, word)
	}

	return strings.Join(jadenWords, " ")
}

func main() {
	fmt.Println(ToJadenCase("How can mirrors be real if our eyes aren't real"))
	fmt.Println(ToJadenCase("When I die. then you will realize"))
}
