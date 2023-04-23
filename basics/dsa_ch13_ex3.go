package main

import (
	"fmt"
	"sort"
)

// 1) O(N^2)

// func FindGreatestNumber(arr []int) int {
// 	for i := 0; i < len(arr); i++ {
// 		isGreatestNumber := true

// 		for j := 0; j < len(arr); j++ {
// 			if arr[j] > arr[i] {
// 				isGreatestNumber = false
// 			}
// 		}

// 		if isGreatestNumber {
// 			return arr[i]
// 		}
// 	}

// 	return 0
// }

// 2) O(N Log N)

func FindGreatestNumber(arr []int) int {
	sort.Ints(arr)
	return arr[len(arr)-1]
}

// 3) O(N)

// func FindGreatestNumber(arr []int) int {
// 	greatestNumber := arr[0]

// 	for _, number := range arr {
// 		if number > greatestNumber {
// 			greatestNumber = number
// 		}
// 	}

// 	return greatestNumber
// }

func main() {
	fmt.Println(FindGreatestNumber([]int{4, 32, 2, 43, 8, 53, 3, 15}))
}
