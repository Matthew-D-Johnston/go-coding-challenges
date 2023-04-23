package main

import (
	"fmt"
	"regexp"
	"strings"
)

func ToCamelCase(s string) string {
	re := regexp.MustCompile(`(-|_)`)
	newS := re.ReplaceAllString(s, " ")
	words := strings.Split(newS, " ")

	camelCaseWords := make([]string, 0)
	camelCaseWords = append(camelCaseWords, words[0])

	for idx, word := range words {
		if idx > 0 {
			chars := strings.Split(string(word), "")
			chars[0] = strings.ToUpper(chars[0])
			camelCaseWords = append(camelCaseWords, strings.Join(chars, ""))
		}
	}

	return strings.Join(camelCaseWords, "")
}

func main() {
	fmt.Println(ToCamelCase("the-stealth-warrior"))
	fmt.Println(ToCamelCase("The_Stealth_Warrior"))
}
