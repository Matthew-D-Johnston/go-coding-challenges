package main

import (
	"fmt"
	"strconv"
)

func CountBits(num uint) int {
	binary := strconv.FormatInt(int64(num), 2)
	sum := 0

	for _, digit := range binary {
		if string(digit) == "1" {
			sum += 1
		}
	}

	return sum
}

func main() {
	fmt.Println(CountBits(1234))
	fmt.Println(CountBits(0))
	fmt.Println(CountBits(4))
	fmt.Println(CountBits(7))
}
