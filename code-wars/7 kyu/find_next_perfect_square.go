package main

import (
	"fmt"
	"math"
)

func FindNextSquare(sq int64) int64 {
	square := float64(sq)

	squareroot := math.Sqrt(square)

	if squareroot != math.Round(squareroot) {
		return -1
	}

	nextSquareroot := squareroot + 1

	nextSquare := math.Pow(nextSquareroot, 2)

	return int64(nextSquare)
}

func main() {
	fmt.Println(FindNextSquare(121))
	fmt.Println(FindNextSquare(625))
	fmt.Println(FindNextSquare(114))
}
