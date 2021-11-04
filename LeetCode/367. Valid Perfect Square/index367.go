package main

import "fmt"

func main() {
	fmt.Println(isPerfectSquare(16))
	fmt.Println(isPerfectSquare(8))
	fmt.Println(isPerfectSquare(86))
	fmt.Println(isPerfectSquare(64))
}

func isPerfectSquare(num int) bool {
	if num == 1 {
		return true
	}
	for i := 2; i <= num/2; i++ {
		if i*i == num {
			return true
		}
	}
	return false
}
