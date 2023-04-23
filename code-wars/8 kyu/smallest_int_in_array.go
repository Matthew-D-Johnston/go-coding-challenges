package main

import "fmt"

func SmallestIntegerFinder(numbers []int) int {
	var min int = numbers[0]

	for _, number := range numbers {
		if number < min {
			min = number
		}
	}

	return min
}

func main() {
	fmt.Println(SmallestIntegerFinder([]int{34, 15, 88, 2}))
	fmt.Println(SmallestIntegerFinder([]int{34, -345, -1, 100}))
}
