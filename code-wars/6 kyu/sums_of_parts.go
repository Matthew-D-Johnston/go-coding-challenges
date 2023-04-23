package main

import "fmt"

// func PartsSums(ls []uint64) []uint64 {
// 	sums := make([]uint64, 0)

// 	for idxI := 0; idxI < len(ls); idxI++ {
// 		var sum uint64 = 0

// 		for idxJ := idxI; idxJ < len(ls); idxJ++ {
// 			sum += ls[idxJ]
// 		}

// 		sums = append(sums, sum)
// 	}

// 	sums = append(sums, 0)

// 	return sums
// }

func PartsSums(ls []uint64) []uint64 {
	sums := make([]uint64, 0)
	var sum uint64 = 0

	for idx := len(ls) - 1; idx >= 0; idx-- {
		sum += ls[idx]
		sums = append([]uint64{sum}, sums...)
	}

	sums = append(sums, 0)

	return sums
}

func main() {
	fmt.Println(PartsSums([]uint64{0, 1, 3, 6, 10}))
	fmt.Println(PartsSums([]uint64{1, 2, 3, 4, 5, 6}))
}
