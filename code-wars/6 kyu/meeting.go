package main

import (
	"fmt"
	"sort"
	"strings"
)

func Meeting(s string) string {
	sortedFamilies := SortFamilies(s)
	combinedNames := CombineNames(sortedFamilies)

	guestList := ""

	for _, combination := range combinedNames {
		family := strings.Split(combination, ":")
		lastName := family[0]
		firstNames := strings.Split(family[1], ",")

		for _, firstName := range firstNames {
			guestList += "(" + strings.ToUpper(lastName) + ", " + strings.ToUpper(firstName) + ")"
		}
	}

	return guestList
}

func SortFamilies(list string) map[string][]string {
	friends := strings.Split(list, ";")
	friendFamilies := make(map[string][]string)

	for _, friend := range friends {
		friendNames := strings.Split(friend, ":")
		friendFamilies[friendNames[1]] = append(friendFamilies[friendNames[1]], friendNames[0])
	}

	for _, firstNames := range friendFamilies {
		sort.Strings(firstNames)
	}

	return friendFamilies
}

func CombineNames(friendFamilies map[string][]string) []string {
	combinedNames := make([]string, 0)

	for lastName, firstNames := range friendFamilies {
		combination := ""
		combination += lastName + ":" + strings.Join(firstNames, ",")
		combinedNames = append(combinedNames, combination)
	}

	sort.Strings(combinedNames)
	return combinedNames
}

func main() {
	fmt.Println(Meeting("Alexis:Wahl;John:Bell;Victoria:Schwarz;Abba:Dorny;Grace:Meta;Ann:Arno;Madison:STAN;Alex:Cornwell;Lewis:Kern;Megan:Stan;Alex:Korn"))
	fmt.Println(Meeting("Fred:Corwill;Wilfred:Corwill;Barney:Tornbull;Betty:Tornbull;Bjon:Tornbull;Raphael:Corwill;Alfred:Corwill"))
}
