package main

import (
	"fmt"
	"sort"
)

func FindMissingNumber(arr []int) int {
	sort.Ints(arr)

	for i := 0; i < len(arr); i++ {
		if i != arr[i] {
			return i
		}
	}

	return -1
}

func main() {
	fmt.Println(FindMissingNumber([]int{5, 2, 4, 1, 0}))
	fmt.Println(FindMissingNumber([]int{9, 3, 2, 5, 6, 7, 1, 0, 4}))
}
