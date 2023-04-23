package main

import (
	"fmt"
	"sort"
)

func MinMax(arr []int) [2]int {
	sort.Ints(arr)
	min := arr[0]
	max := arr[len(arr)-1]

	return [2]int{min, max}
}

func main() {
	fmt.Println(MinMax([]int{1, 2, 3, 4, 5}))
	fmt.Println(MinMax([]int{2334454, 5}))
	fmt.Println(MinMax([]int{1}))
}
