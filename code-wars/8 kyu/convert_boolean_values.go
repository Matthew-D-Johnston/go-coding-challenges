package main

import "fmt"

func BoolToWord(word bool) string {
	if word {
		return "Yes"
	} else {
		return "No"
	}
}

func main() {
	fmt.Println(BoolToWord(true))
	fmt.Println(BoolToWord(false))
}
