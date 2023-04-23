package main

import "fmt"

func PositiveSum(numbers []int) int {
	var sum int
	for i := 0; i < len(numbers); i++ {
		if numbers[i] >= 0 {
			sum = sum + numbers[i]
		}
	}
	return sum
}

func main() {
	arr := []int{1, -4, 7, 12}
	fmt.Println(PositiveSum(arr))
	arr1 := []int{3, 9, 25, -3, -5, 2}
	fmt.Println(PositiveSum((arr1)))
}
