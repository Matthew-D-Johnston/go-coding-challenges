package main

import "fmt"

func GreatestProduct(arr []int) int {
	greatest := arr[0]
	secondGreatest := arr[1]
	lowest := arr[0]
	secondLowest := arr[1]

	for _, number := range arr {
		if number > greatest {
			secondGreatest = greatest
			greatest = number
		}

		if number < greatest && number > secondGreatest {
			secondGreatest = number
		}

		if number < lowest {
			secondLowest = lowest
			lowest = number
		}

		if number > lowest && number < secondLowest {
			secondLowest = number
		}
	}

	product1 := greatest * secondGreatest
	product2 := lowest * secondLowest

	if product1 > product2 {
		return product1
	} else {
		return product2
	}
}

func main() {
	fmt.Println(GreatestProduct([]int{5, -10, -6, 9, 4}))
	fmt.Println(GreatestProduct([]int{-3, 0, 7, -9, 6, -4}))
}
