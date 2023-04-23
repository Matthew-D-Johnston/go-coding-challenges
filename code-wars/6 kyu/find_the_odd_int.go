package main

import "fmt"

func CountValues(seq []int) map[int]int {
	valueCount := make(map[int]int)

	for _, num := range seq {
		valueCount[num] += 1
	}

	return valueCount
}

func FindOdd(seq []int) int {
	valueCount := CountValues(seq)

	for key, value := range valueCount {
		if value%2 == 1 {
			return key
		}
	}

	return 0
}

func main() {
	fmt.Println(FindOdd([]int{1, 2, 2, 3, 3, 3, 4, 3, 3, 3, 2, 2, 1}))
	fmt.Println(FindOdd([]int{20, 1, -1, 2, -2, 3, 3, 5, 5, 1, 2, 4, 20, 4, -1, -2, 5}))
}
