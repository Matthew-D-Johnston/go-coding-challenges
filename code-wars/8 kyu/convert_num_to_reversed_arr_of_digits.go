package main

import (
	"fmt"
	"math"
)

func Digitize(n int) []int {
	if n == 0 {
		return []int{0}
	}

	var reversed []int
	nfloat := float64(n)

	for nfloat > 0 {
		nint := int(nfloat)
		remainder := nint % 10

		reversed = append(reversed, remainder)

		nfloat = math.Floor(nfloat / 10)
	}

	return reversed
}

func main() {
	fmt.Println(Digitize(35231))
	fmt.Println(Digitize(0))
}
