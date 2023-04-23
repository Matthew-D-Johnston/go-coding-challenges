package main

import "fmt"

func CountSheeps(numbers []bool) int {
	var sheep int

	for _, present := range numbers {
		if present {
			sheep += 1
		}
	}

	return sheep
}

func main() {
	fmt.Println(CountSheeps([]bool{true, true, true, false,
		true, true, true, true,
		true, false, true, false,
		true, false, false, true,
		true, true, true, true,
		false, false, true, true}))
}
