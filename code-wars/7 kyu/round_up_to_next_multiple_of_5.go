package main

import "fmt"

func RoundToNext5(n int) int {
	remainder := n % 5

	if remainder == 0 {
		return n
	}

	if n < 0 {
		return n - remainder
	}

	return n - remainder + 5
}

func main() {
	fmt.Println(RoundToNext5(0))
	fmt.Println(RoundToNext5(2))
	fmt.Println(RoundToNext5(3))
	fmt.Println(RoundToNext5(12))
	fmt.Println(RoundToNext5(21))
	fmt.Println(RoundToNext5(30))
	fmt.Println(RoundToNext5(-2))
	fmt.Println(RoundToNext5(-5))
}
