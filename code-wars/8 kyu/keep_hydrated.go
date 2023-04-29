package main

import (
	"fmt"
	"math"
)

func Litres(time float64) int {
	var litres float64 = time * 0.5
	return int(math.Floor(litres))
}

func main() {
	fmt.Println(Litres(3))
	fmt.Println(Litres(6.7))
	fmt.Println(Litres(11.8))
}
