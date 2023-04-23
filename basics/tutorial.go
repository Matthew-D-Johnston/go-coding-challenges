package main

import "fmt"

func main() {
	name := "tim"

	fmt.Println("Before if")
	if name != "tim" {
		fmt.Println("inside if")
	}
	fmt.Println("After if")
}
