package main

import "fmt"

func Solution(str string) []string {
	if len(str)%2 == 1 {
		str += "_"
	}

	charPairs := make([]string, 0)
	for i := 0; i <= len(str)-2; i += 2 {
		charPair := string(str[i : i+2])
		charPairs = append(charPairs, charPair)
	}

	return charPairs
}

func main() {
	fmt.Println(Solution("abc"))
	fmt.Println(Solution("abcdef"))
}
