package main

import "fmt"

func main() {
	fmt.Println(hammingDistance(1, 4))      // 2
	fmt.Println(hammingDistance(1, 3))      // 1
	fmt.Println(hammingDistance(100, 3354)) // 1
}

func hammingDistance(x int, y int) int {
	t := x ^ y
	result := 0
	for {
		if t != 0 {
			bit := t & 1
			if bit == 1 {
				result++
			}
			t = t >> 1
		} else {
			break
		}
	}
	return result
}
