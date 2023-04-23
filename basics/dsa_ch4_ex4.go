package main

import "fmt"

// Greatest number function: O(N^2)--quadratic time

// func GreatestNumber(arr []int) int {
// 	for _, i := range arr {
// 		isIValTheGreatest := true

// 		for _, j := range arr {
// 			if j > i {
// 				isIValTheGreatest = false
// 			}
// 		}
// 		if isIValTheGreatest {
// 			return i
// 		}
// 	}

// 	return 0
// }

// Greatest number function: O(N)--linear time

func GreatestNumber(arr []int) int {
	greatestNum := arr[0]

	for _, i := range arr {
		if i > greatestNum {
			greatestNum = i
		}
	}

	return greatestNum
}

func main() {
	fmt.Println(GreatestNumber([]int{3, 8, 2, 9, 4, 5}))
}
