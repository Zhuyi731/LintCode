package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(fractionToDecimal(1, 90))
	fmt.Println(fractionToDecimal(1, 6))
	fmt.Println(fractionToDecimal(2, 3))
	fmt.Println(fractionToDecimal(4, 333))
	fmt.Println(fractionToDecimal(-4, 333))
	fmt.Println(fractionToDecimal(-4, 33333))
	fmt.Println(fractionToDecimal(1, 2))
	fmt.Println(fractionToDecimal(2, -3))
	fmt.Println(fractionToDecimal(2, 1))
	fmt.Println(fractionToDecimal(7, 9))
}
func fractionToDecimal(numerator, denominator int) string {
	if numerator%denominator == 0 {
		return strconv.Itoa(numerator / denominator)
	}

	s := []byte{}
	if numerator < 0 != (denominator < 0) {
		s = append(s, '-')
	}

	// 整数部分
	numerator = abs(numerator)
	denominator = abs(denominator)
	integerPart := numerator / denominator
	s = append(s, strconv.Itoa(integerPart)...)
	s = append(s, '.')

	// 小数部分
	indexMap := map[int]int{}
	remainder := numerator % denominator
	for remainder != 0 && indexMap[remainder] == 0 {
		indexMap[remainder] = len(s)
		remainder *= 10
		s = append(s, '0'+byte(remainder/denominator))
		remainder %= denominator
	}
	if remainder > 0 { // 有循环节
		insertIndex := indexMap[remainder]
		s = append(s[:insertIndex], append([]byte{'('}, s[insertIndex:]...)...)
		s = append(s, ')')
	}

	return string(s)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
