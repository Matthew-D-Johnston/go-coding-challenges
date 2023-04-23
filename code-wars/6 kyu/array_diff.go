package main

import "fmt"

func ContainsInt(numbers []int, n int) bool {
	for _, num := range numbers {
		if num == n {
			return true
		}
	}

	return false
}

func ArrayDiff(a, b []int) []int {
	differenceArray := make([]int, 0)

	for _, num := range a {
		if !ContainsInt(b, num) {
			differenceArray = append(differenceArray, num)
		}
	}

	return differenceArray
}

func main() {
	fmt.Println(ArrayDiff([]int{1, 2}, []int{1}))
	fmt.Println(ArrayDiff([]int{1, 2, 2}, []int{1}))
	fmt.Println(ArrayDiff([]int{1, 2, 2}, []int{2}))
}
