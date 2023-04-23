package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Revrot(s string, n int) string {
	if n <= 0 || n > len(s) || len(s) == 0 {
		return ""
	}

	remainder := len(s) % n
	mainString := s

	if remainder != 0 {
		mainString = s[0 : len(s)-remainder]
	}

	chunks := make([]string, 0)

	for i := 0; i < len(mainString); i += n {
		chunk := mainString[i : i+n]
		chunks = append(chunks, chunk)
	}

	modifiedChunks := make([]string, 0)

	for _, chunk := range chunks {
		if SumOfDigitsCubed(chunk)%2 == 0 {
			modifiedChunks = append(modifiedChunks, Reverse(chunk))
		} else {
			modifiedChunks = append(modifiedChunks, Rotate(chunk))
		}
	}

	return strings.Join(modifiedChunks, "")
}

func SumOfDigitsCubed(chunk string) int {
	digits := strings.Split(chunk, "")
	sum := 0

	for _, digit := range digits {
		numericDigit, err := strconv.Atoi(digit)

		if err != nil {
			fmt.Println("Conversion error.")
		}

		sum += numericDigit * numericDigit * numericDigit
	}

	return sum
}

func Reverse(chunk string) string {
	reversedChunk := ""

	for i := len(chunk) - 1; i >= 0; i-- {
		reversedChunk += string(chunk[i])
	}

	return reversedChunk
}

func Rotate(chunk string) string {
	return chunk[1:] + string(chunk[0])
}

func main() {
	fmt.Println(Revrot("123456987654", 6))
	fmt.Println(Revrot("66443875", 4))
	fmt.Println(Revrot("66443875", 8))
	fmt.Println(Revrot("", 8))
	fmt.Println(Revrot("123456779", 0))
}
