package main

import (
	"fmt"
	"strings"
)

func TowerBuilder(nFloors int) []string {
	tower := make([]string, 0)
	floorWidth := 2*nFloors - 1

	for i := 1; i <= nFloors; i++ {
		stars := (2 * i) - 1
		spaces := floorWidth - stars
		sideSpaces := spaces / 2
		floor := strings.Repeat(" ", sideSpaces) + strings.Repeat("*", stars) + strings.Repeat(" ", sideSpaces)
		tower = append(tower, floor)
	}

	return tower
}

func main() {
	fmt.Println(TowerBuilder(1))
	fmt.Println(TowerBuilder(2))
	fmt.Println(TowerBuilder(3))
}
