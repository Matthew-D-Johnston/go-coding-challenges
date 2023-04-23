package main

import "fmt"

func LongestConsecutiveSequence(arr []int) int {
	values := make(map[int]bool)
	smallestValue := arr[0]
	largestValue := arr[0]

	for _, value := range arr {
		values[value] = true

		if value < smallestValue {
			smallestValue = value
		}

		if value > largestValue {
			largestValue = value
		}
	}

	currentStreak := 0
	longestStreak := 0

	for value := smallestValue; value <= largestValue; value++ {
		if values[value] {
			currentStreak += 1
			if currentStreak > longestStreak {
				longestStreak = currentStreak
			}
		} else {
			currentStreak = 0
		}
	}

	return longestStreak
}

func main() {
	fmt.Println(LongestConsecutiveSequence(([]int{10, 5, 12, 3, 55, 30, 4, 11, 2})))
	fmt.Println(LongestConsecutiveSequence(([]int{19, 13, 15, 12, 18, 14, 17, 11})))
}
