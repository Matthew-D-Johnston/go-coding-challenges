package main

import (
	"fmt"
	"regexp"
	"sort"
)

func InArray(array1 []string, array2 []string) []string {
	result := make([]string, 0)

	for _, subStr := range array1 {
		containsSubStr := false

		for _, str := range array2 {
			matched, err := regexp.MatchString(subStr, str)

			if err != nil {
				fmt.Println("Matching error.")
			}

			if matched {
				containsSubStr = true
			}
		}

		if containsSubStr && !Contains(result, subStr) {
			result = append(result, subStr)
		}
	}

	sort.Strings(result)
	return result
}

func Contains(arr []string, str string) bool {
	for _, el := range arr {
		if el == str {
			return true
		}
	}

	return false
}

func main() {
	var a1 = []string{"live", "arp", "strong"}
	var a2 = []string{"lively", "alive", "harp", "sharp", "armstrong"}
	fmt.Println(InArray(a1, a2))
	a1 = []string{"cod", "code", "wars", "ewar", "ar"}
	a2 = []string{}
	fmt.Println(InArray(a1, a2))
	a1 = []string{"tarp", "mice", "bull", "live", "arm"}
	a2 = []string{"lively", "alive", "harp", "sharp", "armstrong"}
	fmt.Println(InArray(a1, a2))

}
