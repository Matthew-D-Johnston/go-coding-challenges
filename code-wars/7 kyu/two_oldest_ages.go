package main

import "fmt"

func TwoOldestAges(ages []int) [2]int {
	var oldestAge int
	var secondOldestAge int

	if ages[0] > ages[1] {
		oldestAge = ages[0]
		secondOldestAge = ages[1]
	} else {
		oldestAge = ages[1]
		secondOldestAge = ages[0]
	}

	for idx, age := range ages {
		if idx > 1 {
			if age > oldestAge {
				secondOldestAge = oldestAge
				oldestAge = age
			} else if age > secondOldestAge && age <= oldestAge {
				secondOldestAge = age
			}
		}
	}

	return [2]int{secondOldestAge, oldestAge}
}

func main() {
	var ages1 = []int{1, 2, 10, 8}
	fmt.Println(TwoOldestAges(ages1))
	var ages2 = []int{1, 5, 87, 45, 8, 8}
	fmt.Println(TwoOldestAges(ages2))
	var ages3 = []int{1, 3, 10, 0}
	fmt.Println(TwoOldestAges(ages3))
}
