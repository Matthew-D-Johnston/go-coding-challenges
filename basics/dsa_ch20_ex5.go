package main

import (
	"fmt"
	"strconv"
)

func SortTemperatures(temperatures []float64) []float64 {
	count := make(map[float64]int)

	for _, temperature := range temperatures {
		count[temperature] += 1
	}

	sortedTemperatures := make([]float64, 0)

	for temperature := 97.0; temperature <= 99.0; temperature += 0.1 {
		strRoundedTemperature := fmt.Sprintf("%.1f", temperature)
		roundedTemperature, err := strconv.ParseFloat(strRoundedTemperature, 64)

		if err != nil {
			fmt.Println("Conversion error.")
		}

		fmt.Println(roundedTemperature)

		for i := 1; i <= count[roundedTemperature]; i++ {
			sortedTemperatures = append(sortedTemperatures, roundedTemperature)
		}
	}

	return sortedTemperatures
}

func main() {
	fmt.Println(SortTemperatures([]float64{98.6, 98.0, 97.1, 99.0, 98.9, 97.8, 98.5, 98.2, 98.0, 97.1}))
}
