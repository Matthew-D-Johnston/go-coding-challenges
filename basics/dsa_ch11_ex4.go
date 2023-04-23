package main

import "fmt"

func FindLittleX(str string, index int) int {
	if string(str[0]) == "x" {
		return index
	}

	return FindLittleX(str[1:], index+1)
}

func main() {
	fmt.Println(FindLittleX("abcdefghijklmnopqrstuvwxyz", 0))
}
