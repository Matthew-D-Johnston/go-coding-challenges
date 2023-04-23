package main

import "fmt"

func TotalCharacters(arr []string) int {
	if len(arr) == 1 {
		return len(arr[0])
	}

	return len(arr[0]) + TotalCharacters(arr[1:])
}

func main() {
	fmt.Println(TotalCharacters([]string{"ab", "c", "def", "ghij"}))
}
