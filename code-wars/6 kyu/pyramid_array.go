package main

import "fmt"

func Pyramid(n int) [][]int {
	pyr := make([][]int, 0)
	level := 1

	for level <= n {
		pyrLevel := make([]int, 0)
		ones := 1

		for ones <= level {
			pyrLevel = append(pyrLevel, 1)
			ones += 1
		}

		pyr = append(pyr, pyrLevel)
		level += 1
	}

	return pyr
}

func main() {
	fmt.Println(Pyramid(0))
	fmt.Println(Pyramid(1))
	fmt.Println(Pyramid(2))
	fmt.Println(Pyramid(3))
}
