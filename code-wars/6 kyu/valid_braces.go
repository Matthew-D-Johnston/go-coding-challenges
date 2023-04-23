package main

import (
	"fmt"
	"regexp"
)

func ValidBraces(str string) bool {
	if len(str)%2 == 1 {
		return false
	}

	re1 := regexp.MustCompile(`\((}|])`)
	re2 := regexp.MustCompile(`\[(\)|})`)
	re3 := regexp.MustCompile(`{(\)|])`)

	if re1.MatchString(str) || re2.MatchString(str) || re3.MatchString(str) {
		return false
	}

	re6 := regexp.MustCompile(`\(`)
	re7 := regexp.MustCompile(`\)`)

	if len(re6.FindAllString(str, -1)) != len(re7.FindAllString(str, -1)) {
		return false
	}

	re8 := regexp.MustCompile(`\[`)
	re9 := regexp.MustCompile(`\]`)

	if len(re8.FindAllString(str, -1)) != len(re9.FindAllString(str, -1)) {
		return false
	}

	re10 := regexp.MustCompile(`\{`)
	re11 := regexp.MustCompile(`\}`)

	if len(re10.FindAllString(str, -1)) != len(re11.FindAllString(str, -1)) {
		return false
	}

	return true
}

func main() {
	fmt.Println(ValidBraces("(){}[]"))
	fmt.Println(ValidBraces("([{}])"))
	fmt.Println(ValidBraces("(}"))
	fmt.Println(ValidBraces("[(])"))
	fmt.Println(ValidBraces("[({})](]"))
}
