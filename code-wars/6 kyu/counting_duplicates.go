package main

import (
	"fmt"
	"strings"
)

func CountChars(seq string) map[string]int {
	charCount := make(map[string]int)

	for _, char := range seq {
		lowerCaseChar := strings.ToLower(string(char))
		charCount[lowerCaseChar] += 1
	}

	return charCount
}

func duplicate_count(s1 string) int {
	charCount := CountChars(s1)
	duplicates := 0

	for _, value := range charCount {
		if value >= 2 {
			duplicates += 1
		}
	}

	return duplicates
}

func main() {
	fmt.Println(duplicate_count("aabBcde"))
	fmt.Println(duplicate_count("aA11"))
	fmt.Println(duplicate_count("abcdeaB11"))
	fmt.Println(duplicate_count("indivisibility"))
}
