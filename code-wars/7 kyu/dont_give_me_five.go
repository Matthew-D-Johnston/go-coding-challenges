package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func DontGiveMeFive(start int, end int) int {
	nums := make([]int, 0)

	for i := start; i <= end; i++ {
		matched, _ := regexp.MatchString(`5`, strconv.Itoa(i))

		if matched {
			continue
		}

		nums = append(nums, i)
	}

	return len(nums)
}

func main() {
	fmt.Println(DontGiveMeFive(1, 9))
	fmt.Println(DontGiveMeFive(4, 17))
}
