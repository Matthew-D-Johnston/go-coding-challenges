package main

import "fmt"

func NoSpace(word string) string {
	var result string

	for _, char := range word {
		if string(char) != " " {
			result += string(char)
		}
	}

	return result
}

func main() {
	fmt.Println(NoSpace("hello world"))
}
