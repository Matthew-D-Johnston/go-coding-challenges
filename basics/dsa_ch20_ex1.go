package main

import "fmt"

func TwoSports(arr1 []map[string]string, arr2 []map[string]string) []string {
	names := make(map[string]bool)

	for _, player := range arr1 {
		name := player["first_name"] + " " + player["last_name"]
		names[name] = true
	}

	twoSportsPlayers := make([]string, 0)

	for _, player := range arr2 {
		name := player["first_name"] + " " + player["last_name"]

		if names[name] {
			twoSportsPlayers = append(twoSportsPlayers, name)
		}
	}

	return twoSportsPlayers
}

func main() {
	basketballPlayers := []map[string]string{
		{"first_name": "Jill", "last_name": "Huang", "team": "Gators"},
		{"first_name": "Janko", "last_name": "Barton", "team": "Sharks"},
		{"first_name": "Wanda", "last_name": "Vakulskas", "team": "Sharks"},
		{"first_name": "Jill", "last_name": "Moloney", "team": "Gators"},
		{"first_name": "Luuk", "last_name": "Watkins", "team": "Gators"},
	}
	footballPlayers := []map[string]string{
		{"first_name": "Hanzla", "last_name": "Radosti", "team": "32ers"},
		{"first_name": "Tina", "last_name": "Watkins", "team": "Barleycorns"},
		{"first_name": "Alex", "last_name": "Patel", "team": "32ers"},
		{"first_name": "Jill", "last_name": "Huang", "team": "Barleycorns"},
		{"first_name": "Wanda", "last_name": "Vakulskas", "team": "Barleycorns"},
	}
	fmt.Println(TwoSports(basketballPlayers, footballPlayers))
}
