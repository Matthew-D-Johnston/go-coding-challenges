package main

import "fmt"

func RemoveChar(word string) string {

	if len(word) >= 2 {
		var result string = ""

		for idx, char := range word {
			if idx > 0 && idx < len(word)-1 {
				result += string(char)
			}
		}

		return result
	}

	return word
}

func main() {
	fmt.Println(RemoveChar(("world")))
	fmt.Println(RemoveChar("a"))
	fmt.Println(RemoveChar("to"))
}
