package main

import (
	"fmt"
	"strconv"
	"strings"
)

func MultiplicationReducer(digits []string) int {
	product := 1

	for _, digit := range digits {
		intDigit, err := strconv.Atoi(digit)

		if err == nil {
			product *= intDigit
		} else {
			fmt.Println("Conversion error")
		}
	}

	return product
}

func Persistence(n int) int {
	count := 0

	if n < 10 {
		return count
	}

	for n >= 10 {
		count += 1
		strNumber := strconv.Itoa(n)
		digits := strings.Split(strNumber, "")
		n = MultiplicationReducer(digits)
	}

	return count
}

func main() {
	fmt.Println(Persistence(39))
	fmt.Println(Persistence(999))
	fmt.Println(Persistence(25))
	fmt.Println(Persistence(4))
}
