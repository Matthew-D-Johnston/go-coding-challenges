package main

import "fmt"

func LoveFunc(flower1, flower2 int) bool {
	if (flower1%2 == 0 && flower2%2 == 1) || (flower1%2 == 1 && flower2%2 == 0) {
		return true
	}

	return false
}

func main() {
	fmt.Println(LoveFunc(1, 4))
	fmt.Println(LoveFunc(2, 2))
	fmt.Println(LoveFunc(3, 7))
}
