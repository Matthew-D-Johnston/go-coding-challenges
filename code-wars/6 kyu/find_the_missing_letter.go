package main

import "fmt"

func FindMissingLetter(chars []rune) (missing rune) {

	for idx, char := range chars {
		if idx > 0 {
			previousCharAscii := int(chars[idx-1])
			currentCharAscii := int(char)

			if previousCharAscii+1 != currentCharAscii {
				missing = rune(previousCharAscii + 1)
			}
		}
	}

	return missing
}

func main() {
	fmt.Println(FindMissingLetter([]rune{'a', 'b', 'c', 'd', 'f'}))
	fmt.Println(FindMissingLetter([]rune{'O', 'Q', 'R', 'S'}))
}
