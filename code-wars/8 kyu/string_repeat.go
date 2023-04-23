package main

import "fmt"

func RepeatStr(repetitions int, value string) string {
	var result string
	for reps := 1; reps <= repetitions; reps++ {
		result += value
	}

	return result
}

func main() {
	fmt.Println(RepeatStr(6, "I"))
	fmt.Println(RepeatStr(5, "Hello"))
}
