package main

import "fmt"

func DuplicateValue(arr []string) string {
	hashTable := make(map[string]bool)

	for _, char := range arr {
		if hashTable[char] {
			return char
		} else {
			hashTable[char] = true
		}
	}

	return ""
}

func main() {
	fmt.Println(DuplicateValue([]string{"a", "b", "c", "d", "c", "e"}))
	fmt.Println(DuplicateValue([]string{"a", "b", "c", "d", "e", "a"}))
}
