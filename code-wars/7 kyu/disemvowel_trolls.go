package main

import (
	"fmt"
	"regexp"
)

func Disemvowel(comment string) string {
	re := regexp.MustCompile(`[aeiouAEIOU]`)
	return re.ReplaceAllString(comment, "")
}

func main() {
	fmt.Println(Disemvowel("This website is for losers LOL!"))
}
