package main

import "fmt"

func Divisors(n int) int {
	divisors := make([]int, 0)

	for i := 1; i <= n/2; i++ {
		if n%i == 0 {
			divisors = append(divisors, i)
		}
	}

	divisors = append(divisors, n)

	return len(divisors)
}

func main() {
	fmt.Println(Divisors(4))
	fmt.Println(Divisors(5))
	fmt.Println(Divisors(12))
	fmt.Println(Divisors(30))
}
