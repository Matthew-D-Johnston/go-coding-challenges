package main

import "fmt"

func sumReducer(arr []int) int {
	sum := 0
	for _, num := range arr {
		sum += num
	}

	return sum
}

func FindEvenIndex(arr []int) int {
	for idx := range arr {
		leftSideSlice := arr[:idx]
		rightSideSlice := arr[idx+1:]
		leftSideSum := sumReducer(leftSideSlice)
		rightSideSum := sumReducer(rightSideSlice)

		if leftSideSum == rightSideSum {
			return idx
		}
	}

	return -1
}

func main() {
	fmt.Println(FindEvenIndex([]int{1, 2, 3, 4, 3, 2, 1}))
	fmt.Println(FindEvenIndex([]int{1, 100, 50, -51, 1, 1}))
	fmt.Println(FindEvenIndex([]int{1, 2, 3, 4, 5, 6}))
}
