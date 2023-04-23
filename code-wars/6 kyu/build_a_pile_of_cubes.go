package main

import (
	"fmt"
	"math"
)

func FindNb(m int) int {
	result := 0
	for i := 1; i < m; i++ {
		if nCubes(i) == m {
			result = i
			break
		} else if nCubes(i) > m {
			result = -1
			break
		}
	}
	return result
}

func nCubes(n int) int {
	if n == 0 {
		return 0
	}

	return int(math.Pow(float64(n), 3)) + nCubes(n-1)
}

func main() {
	fmt.Println(FindNb(1071225))
	fmt.Println(FindNb(91716553919377))
	fmt.Println(FindNb(4183059834009))
	fmt.Println(FindNb(9))
}
