package main

import (
	"fmt"
	"strings"
)

func High(s string) string {
	words := strings.Split(s, " ")
	highScore := 0
	var highestScoringWord string

	for _, word := range words {
		score := 0
		for _, char := range word {
			score += int(char) - 96
		}

		if score > highScore {
			highScore = score
			highestScoringWord = word
		}
	}

	return highestScoringWord
}

func main() {
	fmt.Println(High("man i need a taxi up to ubud"))
	fmt.Println(High("what time are we climbing up the volcano"))
	fmt.Println(High("aa b"))
	fmt.Println(High("b aa"))
}
