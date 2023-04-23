package main

import (
	"fmt"
	"regexp"
)

func solution(str, ending string) bool {
	expression := fmt.Sprintf(`%s\z`, ending)
	matched, _ := regexp.MatchString(expression, str)
	return matched
}

func main() {
	fmt.Println(solution("abc", "bc"))
	fmt.Println(solution("abc", "d"))
}
