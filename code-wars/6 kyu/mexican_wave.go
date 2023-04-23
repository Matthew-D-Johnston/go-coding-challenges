package main

import (
	"fmt"
	"strings"
)

func wave(words string) []string {
	waveWords := make([]string, 0)

	for idx, char := range words {
		if string(char) == " " {
			continue
		}
		chars := strings.Split(words, "")
		chars[idx] = strings.ToUpper(string(char))
		waveWords = append(waveWords, strings.Join(chars, ""))
	}

	return waveWords
}

func main() {
	fmt.Println(wave("hello"))
	fmt.Println(wave(" x yz"))
}
