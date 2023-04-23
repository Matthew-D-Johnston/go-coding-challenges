package main

import (
	"fmt"
	"sort"
)

func Comp(array1 []int, array2 []int) bool {
	if len(array1) != len(array2) {
		return false
	}

	array1 = absValueArr(array1)

	sort.Ints(array1)
	sort.Ints(array2)

	for idx := 0; idx < len(array1); idx++ {
		square := array1[idx] * array1[idx]

		if square != array2[idx] {
			return false
		}
	}

	return true
}

func absValueArr(arr []int) []int {
	result := make([]int, 0)

	for _, val := range arr {
		if val < 0 {
			result = append(result, val*-1)
		} else {
			result = append(result, val)
		}
	}

	return result
}

func main() {
	var a1 = []int{121, 144, 19, 161, 19, 144, 19, 11}
	var a2 = []int{11 * 11, 121 * 121, 144 * 144, 19 * 19, 161 * 161, 19 * 19, 144 * 144, 19 * 19}
	fmt.Println(Comp(a1, a2))
	a1 = []int{121, 144, 19, 161, 19, 144, 19, 11}
	a2 = []int{11 * 21, 121 * 121, 144 * 144, 19 * 19, 161 * 161, 19 * 19, 144 * 144, 19 * 19}
	fmt.Println(Comp(a1, a2))
	a1 = nil
	a2 = []int{11 * 11, 121 * 121, 144 * 144, 19 * 19, 161 * 161, 19 * 19, 144 * 144, 19 * 19}
	fmt.Println(Comp(a1, a2))
}
