package main

import "fmt"

func Intersection(arr1 []int, arr2 []int) []int {
	hashTable := make(map[int]bool)

	for _, value := range arr1 {
		hashTable[value] = true
	}

	intersection := make([]int, 0)

	for _, value := range arr2 {
		if hashTable[value] {
			intersection = append(intersection, value)
		}
	}

	return intersection
}

func main() {
	a := []int{1, 2, 3, 4, 5}
	b := []int{0, 2, 4, 6, 8}
	fmt.Println(Intersection(a, b))
	a = []int{1, 2, 3}
	b = []int{0, 2, 3, 5, 8, 9}
	fmt.Println(Intersection(a, b))
	a = []int{4, 5, 6, 7, 8, 9, 10}
	b = []int{3, 5, 7, 9}
	fmt.Println(Intersection(a, b))
}
