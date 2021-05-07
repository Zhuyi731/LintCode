package main

import "fmt"

func main() {
	fmt.Println(xorOperation(5, 0))
	fmt.Println(xorOperation(4, 3))
	fmt.Println(xorOperation(1, 7))
	fmt.Println(xorOperation(10, 5))
}

func xorOperation(n int, start int) int {
	result := start
	for i := start + 2; i < start+n*2; i = i + 2 {
		result = result ^ i
	}
	return result
}
