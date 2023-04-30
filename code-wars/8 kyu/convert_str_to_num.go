package main

import (
	"fmt"
	"strconv"
)

func StringToNumber(str string) int {
	n, _ := strconv.Atoi(str)
	return n
}

func main() {
	fmt.Println(StringToNumber("1234"))
	fmt.Println(StringToNumber("605"))
	fmt.Println(StringToNumber("1405"))
	fmt.Println(StringToNumber("-7"))
}
