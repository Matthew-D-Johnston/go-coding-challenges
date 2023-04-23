package main

import (
	"fmt"
	"sort"
)

func GreatestProductOfThreeNumbers(arr []int) int {
	sort.Ints(arr)
	length := len(arr)
	return arr[length-1] * arr[length-2] * arr[length-3]
}

func main() {
	fmt.Println(GreatestProductOfThreeNumbers([]int{5, 2, 6, 10, 1}))
}
