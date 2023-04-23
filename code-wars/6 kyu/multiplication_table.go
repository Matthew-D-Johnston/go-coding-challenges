package main

import "fmt"

func MultiplicationTable(size int) [][]int {
	table := make([][]int, 0)

	for i := 1; i <= size; i++ {
		row := make([]int, 0)

		for j := 1; j <= size; j++ {
			row = append(row, i*j)
		}

		table = append(table, row)
	}

	return table
}

func main() {
	fmt.Println(MultiplicationTable(3))
	fmt.Println(MultiplicationTable(4))
}
