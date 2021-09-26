package main

import "fmt"

func main() {
	fmt.Println(getSum(2, 3))
	fmt.Println(getSum(0, 1))
	fmt.Println(getSum(0, 0))
	fmt.Println(getSum(-3, 10))
	fmt.Println(getSum(-6, -7))
}

func getSum(a int, b int) int {
	carry := 0
	for b != 0 {
		carry = a & b << 1
		a = a ^ b
		b = carry
	}
	return a
}
