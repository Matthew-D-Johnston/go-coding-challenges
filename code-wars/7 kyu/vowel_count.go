package main

import (
	"fmt"
	"strings"
)

func GetCount(str string) (count int) {
	count = 0
	vowels := [5]string{"a", "e", "i", "o", "u"}

	for _, vowel := range vowels {
		count += strings.Count(str, vowel)
	}

	return count
}

func main() {
	fmt.Println(GetCount("abracadabra"))
	fmt.Println(GetCount("yellow"))
	fmt.Println(GetCount("ingrained"))
}
