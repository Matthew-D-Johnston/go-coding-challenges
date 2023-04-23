package main

import "fmt"

func MissingLetter(str string) string {
	hashTable := make(map[string]bool)

	for i := 0; i < len(str); i++ {
		hashTable[string(str[i])] = true
	}

	alphabet := []string{"a", "b", "c", "d", "e", "f", "g",
		"h", "i", "j", "k", "l", "m", "n",
		"o", "p", "q", "r", "s", "t", "u",
		"v", "w", "x", "y", "z"}

	for _, letter := range alphabet {
		if !hashTable[letter] {
			return letter
		}
	}

	return ""
}

func main() {
	fmt.Println(MissingLetter("the quick brown box jumps over a lazy dog"))
}
