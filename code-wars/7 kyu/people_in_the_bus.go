package main

import "fmt"

func Number(busStops [][2]int) int {
	var people int
	for _, busStop := range busStops {
		people += busStop[0]
		people -= busStop[1]
	}

	return people
}

func main() {
	fmt.Println(Number([][2]int{{10, 0}, {3, 5}, {5, 8}}))
	fmt.Println(Number([][2]int{{3, 0}, {9, 1}, {4, 10}, {12, 2}, {6, 1}, {7, 10}}))
}
