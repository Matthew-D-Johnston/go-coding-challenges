package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func DigPow(n, p int) int {
	nString := strconv.Itoa(n)
	digits := strings.Split(nString, "")

	var sum float64 = 0
	for _, digit := range digits {
		numericDigit, err := strconv.Atoi(digit)
		if err != nil {
			fmt.Println("Conversion failed.")
		}
		sum += math.Pow(float64(numericDigit), float64(p))
		p++
	}

	k := sum / float64(n)

	if k == math.Round(k) {
		return int(k)
	}

	return -1
}

func main() {
	fmt.Println(DigPow(89, 1))
	fmt.Println(DigPow(695, 2))
	fmt.Println(DigPow(46288, 3))
	fmt.Println(DigPow(92, 1))
}
