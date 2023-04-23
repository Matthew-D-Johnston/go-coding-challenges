package main

import "fmt"

func BestTrade(prices []int) int {
	buy := prices[0]
	sell := prices[0]
	greatestProfit := 0

	for _, price := range prices {
		if price < buy {
			buy = price
			sell = price
		} else if price > sell {
			sell = price
		}

		profit := sell - buy

		if profit > greatestProfit {
			greatestProfit = profit
		}
	}

	return greatestProfit
}

func main() {
	fmt.Println(BestTrade([]int{10, 7, 5, 8, 11, 2, 6}))
}
