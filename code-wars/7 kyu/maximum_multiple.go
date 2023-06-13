package main

import "fmt"

func MaxMultiple(d, b int) int {
	n := b
	for n > 0 {
		if n%d == 0 {
			return n
		}

		n -= 1
	}

	return n
}

func main() {
	fmt.Println(MaxMultiple(2, 7))
}
