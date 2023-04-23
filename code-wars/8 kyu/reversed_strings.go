package main

import "fmt"

func Solution(word string) string {
	var reversed string

	for _, char := range word {
		reversed = string(char) + reversed
	}

	return reversed
}

func main() {
	fmt.Println(Solution("world"))
	fmt.Println(Solution("word"))
}
