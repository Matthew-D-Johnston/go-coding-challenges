package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Order(sentence string) string {
	if sentence == "" {
		return ""
	}

	words := strings.Split(sentence, " ")
	numberOfWords := len(words)
	sortedWords := make([]string, numberOfWords, numberOfWords)

	for _, word := range words {
		re := regexp.MustCompile(`\d`)
		digit := re.FindString(word)
		index, err := strconv.Atoi(digit)

		if err != nil {
			fmt.Println("Error: index is not a proper digit")
		}

		sortedWords[index-1] = word
	}

	return strings.Join(sortedWords, " ")
}

func main() {
	fmt.Println(Order("is2 Thi1s T4est 3a"))
	fmt.Println(Order("4of Fo1r pe6ople g3ood th5e the2"))
}
