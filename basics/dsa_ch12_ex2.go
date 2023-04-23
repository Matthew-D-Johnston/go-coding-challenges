package main

import "fmt"

// func Golomb(n int) int {
// 	if n == 1 {
// 		return 1
// 	}

// 	return 1 + Golomb(n-Golomb(Golomb(n-1)))
// }

// Optimization

func Golomb(n int, memo map[int]int) int {
	if n == 1 {
		return 1
	}

	val, ok := memo[n]

	if !ok {
		memo[n] = 1 + Golomb(n-Golomb(Golomb(n-1, memo), memo), memo)
		val = memo[n]
	}

	return val
}

func main() {
	// fmt.Println(Golomb(10))
	// fmt.Println(Golomb(11))
	// fmt.Println(Golomb(15))
	// fmt.Println(Golomb(20))

	fmt.Println(Golomb(10, make(map[int]int)))
	fmt.Println(Golomb(11, make(map[int]int)))
	fmt.Println(Golomb(15, make(map[int]int)))
	fmt.Println(Golomb(20, make(map[int]int)))
}
