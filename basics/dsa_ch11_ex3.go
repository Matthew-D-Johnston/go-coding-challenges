package main

import "fmt"

func TriangularNumber(n int) int {
	if n == 1 {
		return 1
	}

	return n + TriangularNumber(n-1)
}

func main() {
	fmt.Println(TriangularNumber(7))
	fmt.Println(TriangularNumber(3))
}
