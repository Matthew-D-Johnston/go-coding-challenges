package main

import (
	"fmt"
)

func NbYear(p0 int, percent float64, aug int, p int) int {
	population := float64(p0)
	var years int

	for population <= float64(p) {
		population += (float64(population) * (percent / 100)) + float64(aug)
		years += 1
	}

	return years
}

func main() {
	fmt.Println(NbYear(1500, 5, 100, 5000))
	fmt.Println(NbYear(1500000, 2.5, 10000, 2000000))
	fmt.Println(NbYear(1500000, 0.25, 1000, 2000000))
	fmt.Println(NbYear(1500000, 0.25, -1000, 2000000))
}
