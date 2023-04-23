package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func PrinterError(s string) string {
	re := regexp.MustCompile(`[n-z]`)
	errorMatches := re.FindAllString(s, -1)
	numerator := strconv.Itoa(len(errorMatches))
	denominator := strconv.Itoa(len(s))
	return numerator + "/" + denominator
}

func main() {
	fmt.Println(PrinterError("aaabbbbhaijjjm"))
	fmt.Println(PrinterError("aaaxbbbbyyhwawiwjjjwwm"))
}
