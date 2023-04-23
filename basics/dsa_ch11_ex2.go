package main

import "fmt"

func EvenNumbers(arr []int) []int {
	if len(arr) == 1 {
		if arr[0]%2 == 0 {
			return arr
		} else {
			return make([]int, 0)
		}
	}

	if arr[0]%2 == 0 {
		return append(EvenNumbers(arr[1:]), arr[0])
	}

	return EvenNumbers(arr[1:])
}

func main() {
	fmt.Println(EvenNumbers([]int{2, 3, 4, 5, 7}))
	fmt.Println(EvenNumbers([]int{1, 3, 4, 6, 7, 9, 10}))
	fmt.Println(EvenNumbers([]int{1, 3, 5, 7}))
}
