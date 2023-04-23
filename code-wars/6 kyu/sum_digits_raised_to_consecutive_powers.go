package main

import (
	"fmt"
	"math/big"
	"strings"
)

func SumDigPow(a, b uint64) []uint64 {
	result := make([]uint64, 0)

	for i := a; i <= b; i++ {
		bigIntI := big.NewInt(int64(i))

		if bigIntI.Cmp(SumConsecDigPow(bigIntI)) == 0 {
			result = append(result, i)
		}
	}

	return result
}

func SumConsecDigPow(num *big.Int) *big.Int {
	sum := big.NewInt(0)
	strNum := num.String()
	digits := strings.Split(strNum, "")

	for exponent := big.NewInt(int64(len(digits))); exponent.Cmp(big.NewInt(0)) == 1; exponent.Sub(exponent, big.NewInt(1)) {
		quotient := big.NewInt(0)
		remainder := big.NewInt(0)
		quotient.DivMod(num, big.NewInt(10), remainder)

		calc := big.NewInt(0)
		zero := big.NewInt(0)
		calc.Exp(remainder, exponent, zero)

		sum.Add(sum, calc)
		num = quotient
	}

	return sum
}

// func SumConsecDigPow(num uint64) uint64 {
// 	var sum float64 = 0
// 	strNum := strconv.Itoa(int(num))
// 	digits := strings.Split(strNum, "")

// 	for idx, strDigit := range digits {
// 		// digit, err := strconv.Atoi(strDigit)
// 		digit := ConvertStrToInt(strDigit)
// 		// digit := new(big.Float)
// 		// digit.SetString(strDigit)

// 		// digit, err := strconv.ParseUint(strDigit, 10, 64)

// 		// if err != nil {
// 		// 	fmt.Println("Conversion error.")
// 		// }

// 		sum += math.Pow(float64(digit), float64(idx+1))
// 	}

// 	return uint64(sum)
// }

// func ConvertStrToInt(str string) int {
// 	digits := strings.Split(str, "")
// 	sum := 0
// 	power := 0

// 	for idx := len(digits) - 1; idx >= 0; idx-- {
// 		digit, err := strconv.Atoi(digits[idx])

// 		if err != nil {
// 			fmt.Println("Conversion error.")
// 		}

// 		sum += digit * int(math.Pow(10, float64(power)))
// 		power += 1
// 	}

// 	return sum
// }

func main() {
	// fmt.Println(SumConsecDigPow(89))
	// fmt.Println(SumConsecDigPow(135))
	// fmt.Println(SumConsecDigPow(9))
	// fmt.Println(SumConsecDigPow(10))
	fmt.Println(SumDigPow(1, 10))
	fmt.Println(SumDigPow(1, 100))
	fmt.Println(SumDigPow(90, 100))
	fmt.Println(SumDigPow(89, 135))
	// fmt.Println(SumDigPow(12157692622039623127, 12157692622039625670))
	// digit, err := strconv.ParseUint("12157692622039623127", 10, 64)

	// if err != nil {
	// 	fmt.Println("Conversion error.")
	// } else {
	// 	fmt.Println(digit)
	// }

	// fmt.Println(ConvertStrToInt("12157692622039623127"))
	// sum := big.NewInt(0)
	// calc := big.NewInt(0)
	// remainderBigInt := big.NewInt(int64(5))
	// exponentBigInt := big.NewInt(int64(3))
	// zero := big.NewInt(0)
	// calc.Exp(remainderBigInt, exponentBigInt, zero)
	// fmt.Println(calc)
	// sum.Add(sum, calc)
	// fmt.Println(sum)
}
