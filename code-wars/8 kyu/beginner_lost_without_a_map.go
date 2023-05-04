package main

import (
	"fmt"
)

func Maps(x []int) []int {
	var mapped []int

	for _, value := range x {
		mapped = append(mapped, value*2)
	}

	return mapped
}

func main() {
	fmt.Println(Maps([]int{1, 2, 3}))
}
