package main

import "fmt"

func SequenceSum(start, end, step int) int {
	if start > end {
		return 0
	}

	sum := 0
	current := start

	for current <= end {
		sum += current
		current += step
	}

	return sum
}

func main() {
	fmt.Println(SequenceSum(2, 2, 2))
	fmt.Println(SequenceSum(2, 6, 2))
	fmt.Println(SequenceSum(1, 5, 1))
	fmt.Println(SequenceSum(1, 5, 3))
}
