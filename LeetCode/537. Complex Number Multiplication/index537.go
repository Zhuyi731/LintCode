package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(complexNumberMultiply("1+-1i", "1+-1i"))
	fmt.Println(complexNumberMultiply("1+1i", "1+1i"))
	fmt.Println(complexNumberMultiply("1+1i", "1+-1i"))
}

func complexNumberMultiply(num1 string, num2 string) string {
	num1R := 0
	num1C := 0
	num2R := 0
	num2C := 1

	for i := 0; i < len(num1); i++ {
		if num1[i] == '+' {
			num1R, _ = strconv.Atoi(num1[:i])
			num1C, _ = strconv.Atoi(num1[i+1 : len(num1)-1])
		}
	}

	for i := 0; i < len(num2); i++ {
		if num2[i] == '+' {
			num2R, _ = strconv.Atoi(num2[:i])
			num2C, _ = strconv.Atoi(num2[i+1 : len(num2)-1])
		}
	}

	resultR := num2R*num1R - num2C*num1C
	resultC := num1R*num2C + num1C*num2R

	return strconv.Itoa(resultR) + "+" + strconv.Itoa((resultC)) + "i"
}
