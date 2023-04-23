package main

import (
	"fmt"
	"strconv"
)

func NumberToString(n int) string {
	return strconv.Itoa(n)
}

func main() {
	fmt.Println(NumberToString(123))
	fmt.Println(NumberToString(999))
	fmt.Println(NumberToString(-100))
	fmt.Println(123)
}
