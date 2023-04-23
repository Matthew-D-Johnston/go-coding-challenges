package main

import (
	"fmt"
	"strings"
)

func AbbrevName(name string) string {
	var abbreviation string = strings.ToUpper(string(name[0])) + "."

	for idx, char := range name {
		if string(char) == " " {
			abbreviation += strings.ToUpper(string(name[idx+1]))
		}
	}

	return abbreviation
}

func main() {
	fmt.Println(AbbrevName("Sam Harris"))
	fmt.Println(AbbrevName("patrick feeney"))
}
