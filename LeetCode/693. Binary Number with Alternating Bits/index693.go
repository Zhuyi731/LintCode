package main

import "fmt"

func main() {
	fmt.Println(hasAlternatingBits(3))  // true
	fmt.Println(hasAlternatingBits(1))  // true
	fmt.Println(hasAlternatingBits(2))  // true
	fmt.Println(hasAlternatingBits(42)) // true
	fmt.Println(hasAlternatingBits(5))  // true
	fmt.Println(hasAlternatingBits(7))  // false
	fmt.Println(hasAlternatingBits(11)) // false
	fmt.Println(hasAlternatingBits(21)) // true
	fmt.Println(hasAlternatingBits(22)) // false
}

func hasAlternatingBits(n int) bool {
	prev := n % 2
	n = n / 2
	for {
		cur := n % 2
		n = n / 2
		if cur^prev == 0 {
			return false
		}
		if n == 0 {
			return true
		}
		prev = cur
	}
}
