package main

import "fmt"

func main() {
	fmt.Println(singleNumber([]int{1, 2, 1, 3, 2, 5}))
	fmt.Println(singleNumber([]int{0, -1}))
}

func singleNumber(nums []int) []int {
	xorsum := 0
	for _, v := range nums {
		xorsum ^= v
	}
	min1 := xorsum & (-xorsum)
	result := []int{0, 0}
	for _, v := range nums {
		if v&min1 == 0 {
			result[0] ^= v
		}
	}
	result[1] = xorsum ^ result[0]
	return result
}
