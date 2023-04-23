package main

import "fmt"

func ConcatenateStrings(list []string, k int) []string {
	concatenatedStrings := make([]string, 0)

	for outerIdx := 0; outerIdx <= len(list)-k; outerIdx++ {
		concatenatedStr := ""

		for innerIdx := outerIdx; innerIdx < outerIdx+k; innerIdx++ {
			concatenatedStr += list[innerIdx]
		}

		concatenatedStrings = append(concatenatedStrings, concatenatedStr)
	}

	return concatenatedStrings
}

func LongestString(list []string) string {
	var longestLength int
	var longest string

	for _, str := range list {
		if len(str) > longestLength {
			longestLength = len(str)
			longest = str
		}
	}

	return longest
}

func LongestConsec(strarr []string, k int) string {
	length := len(strarr)

	if length == 0 || k > length || k <= 0 {
		return ""
	}

	concatenatedStrings := ConcatenateStrings(strarr, k)

	return LongestString(concatenatedStrings)
}

func main() {
	fmt.Println(LongestConsec([]string{"zone", "abigail", "theta", "form", "libe", "zas"}, 2))
	fmt.Println(LongestConsec([]string{"ejjjjmmtthh", "zxxuueeg", "aanlljrrrxx", "dqqqaaabbb", "oocccffuucccjjjkkkjyyyeehh"}, 1))
}
