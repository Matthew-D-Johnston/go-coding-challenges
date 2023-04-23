package main

import (
	"fmt"
	"sort"
)

func QueueTime(customers []int, n int) int {
	tills := make(map[int][]int)

	sortedCustomers := make([]int, len(customers))
	copy(sortedCustomers, customers)
	sort.Ints(sortedCustomers)
	longestWaitTime := sortedCustomers[len(sortedCustomers)-1]

	for till := 1; till <= n; till++ {
		if till <= len(customers) {
			tills[till] = []int{customers[till-1]}
		}
	}

	for customer := n; customer < len(customers); customer++ {
		shortestWaitTime := longestWaitTime
		openTill := 0

		for till, waitTimes := range tills {
			waitTime := waitTimes[len(waitTimes)-1]

			if waitTime <= shortestWaitTime {
				shortestWaitTime = waitTime
				openTill = till
			}
		}

		tills[openTill] = append(tills[openTill], customers[customer])
	}

	longestTotalWaitTime := 0

	for _, waitTimes := range tills {
		totalWaitTime := 0

		for _, waitTime := range waitTimes {
			totalWaitTime += waitTime
		}

		if totalWaitTime > longestTotalWaitTime {
			longestTotalWaitTime = totalWaitTime
		}
	}

	return longestTotalWaitTime
}

func main() {
	fmt.Println(QueueTime([]int{1, 2, 3, 4}, 1))
	fmt.Println(QueueTime([]int{2, 2, 3, 3, 4, 4}, 2))
	fmt.Println(QueueTime([]int{1, 2, 3, 4, 5}, 100))
	fmt.Println(QueueTime([]int{10, 2, 3, 3}, 2))
	fmt.Println(QueueTime([]int{5, 3, 4}, 1))
	fmt.Println(QueueTime([]int{2, 3, 10}, 2))
}
