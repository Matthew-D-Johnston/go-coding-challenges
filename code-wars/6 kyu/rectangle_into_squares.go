package main

import (
	"fmt"
	"sort"
)

func SquaresInRect(lng int, wdth int) []int {
	squares := make([]int, 0)

	if lng == wdth {
		return nil
	}

	area := lng * wdth

	for area > 0 {
		dimensions := []int{lng, wdth}
		sort.Ints(dimensions)
		smallestDimension := dimensions[0]
		largestDimension := dimensions[1]

		squares = append(squares, smallestDimension)
		wdth = largestDimension - smallestDimension
		lng = smallestDimension
		area = area - (smallestDimension * smallestDimension)
	}

	return squares
}

func main() {
	fmt.Println(SquaresInRect(5, 3))
	fmt.Println(SquaresInRect(3, 5))
	fmt.Println(SquaresInRect(5, 5))
	fmt.Println(SquaresInRect(20, 14))
}
