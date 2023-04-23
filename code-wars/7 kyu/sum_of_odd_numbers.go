package main

import "fmt"

func oddRows(n int) int {
	if n == 1 {
		return 1
	}

	return 2*(n-1) + oddRows(n-1)
}

func sumByTwos(startInt int, numberOfTwos int) int {
	sum := startInt
	for i := 1; i < numberOfTwos; i++ {
		sum += startInt + (2 * i)
	}

	return sum
}

func RowSumOddNumbers(n int) int {
	firstNumberOfRow := oddRows(n)
	return sumByTwos(firstNumberOfRow, n)
}

func main() {
	fmt.Println(RowSumOddNumbers(1))
	fmt.Println(RowSumOddNumbers(2))
	fmt.Println(RowSumOddNumbers(3))
	fmt.Println(RowSumOddNumbers(4))
}
