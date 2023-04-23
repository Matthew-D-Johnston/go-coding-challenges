package main

import "fmt"

func MakeNegative(x int) int {
	if x < 0 {
		return x
	} else {
		return x * -1
	}
}

func main() {
	fmt.Println(MakeNegative(-5))
	fmt.Println(MakeNegative(-11))
	fmt.Println(MakeNegative(5))
	fmt.Println(MakeNegative(0))
	fmt.Println(MakeNegative(43))
}
