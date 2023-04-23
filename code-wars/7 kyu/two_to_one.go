package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

func TwoToOne(s1 string, s2 string) string {
	result := ""
	for _, char := range s1 {
		matched, _ := regexp.MatchString(string(char), result)
		if !matched {
			result += string(char)
		}
	}

	for _, char := range s2 {
		matched, _ := regexp.MatchString(string(char), result)
		if !matched {
			result += string(char)
		}
	}

	resultSlice := strings.Split(result, "")
	sort.Strings(resultSlice)
	result = strings.Join(resultSlice, "")
	return result
}

func main() {
	fmt.Println(TwoToOne("xyaabbbccccdefww", "xxxxyyyyabklmopq"))
}
