package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maximumGap([]int{3, 6, 9, 1}))
	fmt.Println(maximumGap([]int{100, 3, 2, 1}))
	fmt.Println(maximumGap([]int{10}))
}

func maximumGap(nums []int) int {
	if len(nums) < 2 {
		return 0
	}

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	currentMaxGap := 0

	for i, v := range nums {
		if i == 0 {
			continue
		}

		gap := v - nums[i-1]
		if gap > currentMaxGap {
			currentMaxGap = gap
		}
	}

	return currentMaxGap
}
