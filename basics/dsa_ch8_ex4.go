package main

import "fmt"

func FirstNonDuplicate(str string) string {
	hashTable := make(map[string]int)

	for i := 0; i < len(str); i++ {
		hashTable[string(str[i])] += 1
	}

	for i := 0; i < len(str); i++ {
		if hashTable[string(str[i])] == 1 {
			return string(str[i])
		}
	}

	return ""
}

func main() {
	fmt.Println(FirstNonDuplicate("minimum"))
	fmt.Println(FirstNonDuplicate("alabama"))
}
