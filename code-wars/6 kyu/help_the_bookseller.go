package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// func StockList(listArt []string, listCat []string) string {
// 	if len(listArt) == 0 || len(listCat) == 0 {
// 		return ""
// 	}

// 	catTotals := make(map[string]int)

// 	for _, category := range listCat {
// 		catTotals[category] = 0

// 		for _, code := range listArt {
// 			if string(code[0]) == category {
// 				re := regexp.MustCompile(`\d+`)
// 				match := re.FindAllString(code, -1)
// 				quantity, err := strconv.Atoi(match[0])

// 				if err != nil {
// 					fmt.Println("Conversion error.")
// 				}

// 				catTotals[category] += quantity
// 			}
// 		}
// 	}
// 	catTotalsArray := make([]string, 0)

// 	for category, quantity := range catTotals {
// 		catTotal := fmt.Sprint("(", category, " : ", quantity, ")")
// 		catTotalsArray = append(catTotalsArray, catTotal)
// 	}
// 	return strings.Join(catTotalsArray, " - ")
// }

// Refactored to maintain order of `listCat`

func StockList(listArt []string, listCat []string) string {
	if len(listArt) == 0 || len(listCat) == 0 {
		return ""
	}

	catTotalsArray := make([]string, 0)

	for _, category := range listCat {
		catTotal := 0

		for _, code := range listArt {
			if string(code[0]) == category {
				re := regexp.MustCompile(`\d+`)
				match := re.FindAllString(code, -1)
				quantity, err := strconv.Atoi(match[0])

				if err != nil {
					fmt.Println("Conversion error.")
				}

				catTotal += quantity
			}
		}

		catTotalFormatted := fmt.Sprint("(", category, " : ", catTotal, ")")
		catTotalsArray = append(catTotalsArray, catTotalFormatted)
	}

	return strings.Join(catTotalsArray, " - ")
}

func main() {
	var b = []string{"BBAR 150", "CDXE 515", "BKWR 250", "BTSQ 890", "DRTY 600"}
	var c = []string{"A", "B", "C", "D"}
	fmt.Println(StockList(b, c))
	b = []string{"ABAR 200", "CDXE 500", "BKWR 250", "BTSQ 890", "DRTY 600"}
	c = []string{"A", "B"}
	fmt.Println(StockList(b, c))
}
