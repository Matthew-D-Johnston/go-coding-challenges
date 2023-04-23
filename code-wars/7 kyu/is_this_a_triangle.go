package main

import (
	"fmt"
)

func IsTriangle(a, b, c int) bool {
	switch {
	case a+b < c:
		return false
	case a+c < b:
		return false
	case b+c < a:
		return false
	case a <= 0:
		return false
	case b <= 0:
		return false
	case c <= 0:
		return false
	}

	return true
}

func main() {
	fmt.Println(IsTriangle(5, 1, 2))
	fmt.Println(IsTriangle(1, 5, 2))
	fmt.Println(IsTriangle(2, 5, 1))
	fmt.Println(IsTriangle(4, 2, 3))
}
