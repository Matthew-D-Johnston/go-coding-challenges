package main

import "fmt"

func Reverse(arr []int) []int {
	leftIndex := 0
	rightIndex := len(arr) - 1

	for leftIndex < rightIndex {
		leftValue := arr[leftIndex]
		rightValue := arr[rightIndex]
		arr[leftIndex] = rightValue
		arr[rightIndex] = leftValue
		leftIndex++
		rightIndex--
	}

	return arr
}

func main() {
	fmt.Println(Reverse([]int{1, 2, 3, 4, 5, 6, 7}))
	fmt.Println(Reverse([]int{10, 11, 12, 13, 14, 15}))
}
