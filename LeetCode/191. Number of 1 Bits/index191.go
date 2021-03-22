package main

import "fmt"

func main() {
	fmt.Println(hammingWeight((00000000000000000000000000001011)))
	fmt.Println(hammingWeight((00000000000000000000000010000000)))
	// fmt.Println(hammingWeight((11111111111111111111111111111101)))
}

func hammingWeight(num uint32) int {
	ct := 0
	for num > 0 {
		if num%2 == 1 {
			ct++
		}
		num = num / 2
	}
	return ct
}
