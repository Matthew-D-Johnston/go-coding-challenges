package main

import "fmt"

func SquareSum(numbers []int) int {
	var sum int

	for _, number := range numbers {
		sum += number * number
	}

	return sum
}

func main() {
	fmt.Println(SquareSum([]int{1, 2, 2}))
	fmt.Println(SquareSum([]int{3, 2, 4, 1}))
}
