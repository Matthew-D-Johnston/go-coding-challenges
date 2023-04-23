package main

import "fmt"

func Tribonacci(signature [3]float64, n int) []float64 {
	sequence := make([]float64, 0)
	if n <= 3 {
		switch n {
		case 0:
			return sequence
		case 1:
			return signature[:1]
		case 2:
			return signature[:2]
		case 3:
			return signature[:3]
		}
	}
	sequence = signature[:3]

	for len(sequence) < n {
		var sum float64 = 0

		for i := len(sequence) - 1; i >= len(sequence)-3; i-- {
			sum += sequence[i]
		}

		sequence = append(sequence, sum)
	}

	return sequence
}

func main() {
	fmt.Println(Tribonacci([3]float64{1, 1, 1}, 10))
	fmt.Println(Tribonacci([3]float64{0, 0, 1}, 10))
	fmt.Println(Tribonacci([3]float64{0, 1, 1}, 10))
	fmt.Println(Tribonacci([3]float64{0, 0, 0}, 10))
}
