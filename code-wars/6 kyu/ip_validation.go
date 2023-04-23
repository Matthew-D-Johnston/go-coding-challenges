package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Is_valid_ip(ip string) bool {
	octets := strings.Split(ip, ".")

	if len(octets) != 4 {
		return false
	}

	for _, octet := range octets {
		matched, err := regexp.MatchString(`\A0.`, octet)

		if err != nil {
			fmt.Println("Matching error.")
		}

		if matched {
			return false
		}

		matched, err = regexp.MatchString(`[a-zA-Z ]`, octet)

		if err != nil {
			fmt.Println("Matching error.")
		}

		if matched {
			return false
		}

		number, err := strconv.Atoi(octet)

		if err != nil {
			fmt.Println("Conversion error.")
		}

		if number < 0 || number > 255 {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(Is_valid_ip("1.2.3.4"))
	fmt.Println(Is_valid_ip("123.45.67.89"))
	fmt.Println(Is_valid_ip("1.2.3"))
	fmt.Println(Is_valid_ip("1.2.3.4.5"))
	fmt.Println(Is_valid_ip("123.456.78.90"))
	fmt.Println(Is_valid_ip("123.045.067.089"))
}
