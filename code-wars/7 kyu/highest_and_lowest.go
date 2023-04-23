package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func HighAndLow(in string) string {
	inSlice := strings.Split(in, " ")
	converted := make([]int, 0)

	for _, num := range inSlice {
		integer, _ := strconv.Atoi(string(num))
		converted = append(converted, integer)
	}
	sort.Ints(converted)
	highest := converted[len(inSlice)-1]
	lowest := converted[0]
	return strconv.Itoa(highest) + " " + strconv.Itoa(lowest)
}

func main() {
	fmt.Println(HighAndLow("1 2 3 4 5"))
	fmt.Println(HighAndLow("1 2 -3 4 5"))
	fmt.Println(HighAndLow("1 9 3 4 -5"))
	fmt.Println(HighAndLow("8 3 -5 42 -1 0 0 -9 4 7 4 -4"))
}
