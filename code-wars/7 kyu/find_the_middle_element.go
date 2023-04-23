package main

import (
	"fmt"
	"sort"
)

func Gimme(array [3]int) int {
	arrayCopy := array
	slice := arrayCopy[0:len(array)]
	sort.Ints(slice)
	middleEl := slice[1]
	var idx int

	for i, num := range array {
		if num == middleEl {
			idx = i
		}
	}

	return idx
}

func main() {
	fmt.Println(Gimme([3]int{2, 3, 1}))
	fmt.Println(Gimme([3]int{5, 10, 14}))
	fmt.Println(Gimme([3]int{15, 10, 14}))
}
