package main

import "fmt"

func Summation(n int) int {
	var sum int
	for i := 1; i <= n; i++ {
		sum += i
	}

	return sum
}

func main() {
	fmt.Println(Summation(2))
	fmt.Println(Summation(8))
}
