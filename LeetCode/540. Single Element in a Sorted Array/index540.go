package main

import "fmt"

func main() {
	fmt.Println(singleNonDuplicate([]int{1, 1, 2, 3, 3, 4, 4, 8, 8}))
	fmt.Println(singleNonDuplicate([]int{3, 3, 7, 7, 10, 11, 11}))
}

func singleNonDuplicate(nums []int) int {
	result := 0
	for _, v := range nums {
		result ^= v
	}
	return result
}
