package main

import "fmt"

func main() {
	fmt.Println(isPowerOfFour(16))
	fmt.Println(isPowerOfFour(64))
	fmt.Println(isPowerOfFour(60))
	fmt.Println(isPowerOfFour(1))
	fmt.Println(isPowerOfFour(31))
}

var result map[int]bool

func init() {
	result = map[int]bool{}
	t := 1
	for {
		result[t] = true
		t = t * 2
		if t > (2<<31 - 1) {
			return
		}
	}
}

func isPowerOfTwo(n int) bool {
	return result[n]
}
