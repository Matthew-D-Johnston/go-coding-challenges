package main

import "fmt"

func IsValidWalk(walk []rune) bool {
	if len(walk) != 10 {
		return false
	}

	origin := map[string]int{"x": 0, "y": 0}

	for _, runeDirection := range walk {
		charDirection := string(runeDirection)

		switch charDirection {
		case "e":
			origin["x"] += 1
		case "w":
			origin["x"] -= 1
		case "n":
			origin["y"] += 1
		case "s":
			origin["y"] -= 1
		}
	}

	if origin["x"] == 0 && origin["y"] == 0 {
		return true
	}

	return false
}

func main() {
	fmt.Println(IsValidWalk([]rune{'n', 's', 'n', 's', 'n', 's', 'n', 's', 'n', 's'}))
	fmt.Println(IsValidWalk([]rune{'w', 'e', 'w', 'e', 'w', 'e', 'w', 'e', 'w', 'e', 'w', 'e'}))
	fmt.Println(IsValidWalk([]rune{'n', 'e', 's', 'w', 'n', 'e', 's', 'w', 'w', 'e'}))
}
