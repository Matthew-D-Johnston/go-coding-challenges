package main

import "fmt"

func AddUntil100(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	remainingSum := AddUntil100(arr[1:])

	if arr[0]+remainingSum > 100 {
		return remainingSum
	} else {
		return arr[0] + remainingSum
	}
}

func main() {
	fmt.Println(AddUntil100([]int{20, 40, 50}))
}
