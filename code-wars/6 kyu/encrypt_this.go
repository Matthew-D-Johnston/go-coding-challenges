package main

import (
	"fmt"
	"strconv"
	"strings"
)

func EncryptThis(text string) string {
	words := strings.Split(text, " ")
	encryptedWords := make([]string, 0)

	for _, word := range words {
		encrypted := ""
		ascii := strconv.Itoa(int(word[0]))
		encrypted += ascii

		if len(word) == 2 {
			encrypted += string(word[1])
		} else if len(word) > 2 {
			secondChar := word[1]
			lastChar := word[len(word)-1]
			middleChars := word[2 : len(word)-1]

			encrypted += string(lastChar)
			encrypted += middleChars
			encrypted += string(secondChar)
		}

		encryptedWords = append(encryptedWords, encrypted)
	}

	encryptedText := strings.Join(encryptedWords, " ")
	return encryptedText
}

func main() {
	fmt.Println(EncryptThis("Hello"))
	fmt.Println(EncryptThis("good"))
	fmt.Println(EncryptThis("hello world"))
	fmt.Println(EncryptThis("A wise old owl lived in an oak"))
}
