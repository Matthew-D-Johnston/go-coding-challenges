package main

import "fmt"

func MissingNumber(arr []int) int {
	existingNumbers := make(map[int]bool)

	for _, number := range arr {
		existingNumbers[number] = true
	}

	for i := 0; i < len(arr); i++ {
		if !existingNumbers[i] {
			return i
		}
	}

	return -1
}

func main() {
	fmt.Println(MissingNumber([]int{2, 3, 0, 6, 1, 5}))
	fmt.Println(MissingNumber([]int{8, 2, 3, 9, 4, 7, 5, 0, 6}))
}
