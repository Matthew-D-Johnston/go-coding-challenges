package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func NbDig(n int, d int) int {
	var digitSum int
	re := regexp.MustCompile(fmt.Sprintf("%d", d))

	for i := 0; i <= n; i++ {
		square := i * i
		matches := re.FindAllString(strconv.Itoa(square), -1)
		digitSum += len(matches)
	}

	return digitSum
}

func main() {
	fmt.Println(NbDig(10, 1))
	fmt.Println(NbDig(25, 1))
}
