package main

import (
	"fmt"
	"strconv"
)

// func UniquePaths(rows int, columns int) int {
// 	if rows == 1 || columns == 1 {
// 		return 1
// 	}
// 	return UniquePaths(rows-1, columns) + UniquePaths(rows, columns-1)
// }

// Optimization
func UniquePaths(rows int, columns int, memo map[string]int) int {
	if rows == 1 || columns == 1 {
		return 1
	}

	key := strconv.Itoa(rows) + strconv.Itoa(columns)

	val, ok := memo[key]

	if !ok {
		memo[key] = UniquePaths(rows-1, columns, memo) + UniquePaths(rows, columns-1, memo)
		val = memo[key]
	}
	return val
}

func main() {
	// fmt.Println(UniquePaths(2, 2))
	// fmt.Println(UniquePaths(3, 7))
	// fmt.Println(UniquePaths(5, 8))

	fmt.Println(UniquePaths(2, 2, make(map[string]int)))
	fmt.Println(UniquePaths(3, 7, make(map[string]int)))
	fmt.Println(UniquePaths(5, 8, make(map[string]int)))
}
