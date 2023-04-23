package main

import (
	"fmt"
	"strings"
)

func Accum(s string) string {
	result := ""

	for idx, char := range s {
		result += strings.ToUpper(string(char))
		result += strings.ToLower(strings.Repeat(string(char), idx))

		if idx < len(s)-1 {
			result += "-"
		}
	}

	return result
}

func main() {
	fmt.Println(Accum("abcd"))
	fmt.Println(Accum("RqaEzty"))
	fmt.Println(Accum("cwAt"))
}
