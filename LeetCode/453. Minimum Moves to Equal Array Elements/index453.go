package main

import "fmt"

func main() {
	fmt.Println(minMoves([]int{1, 2, 3}))
	fmt.Println(minMoves([]int{1, 1, 1}))
}

func minMoves(nums []int) int {
	min := 2<<32 - 1
	for _, v := range nums {
		if v < min {
			min = v
		}
	}
	total := 0
	for _, v := range nums {
		total += v - min
	}

	return total
}
