package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minimumDifference([]int{9, 4, 1, 7}, 2))
}

func minimumDifference(nums []int, k int) int {
	if k == 1 {
		return 0
	}

	sort.Slice(nums, func(i, j int) bool {
		return nums[i]-nums[j] >= 0
	})

	min := 0xfffffff
	for i := 0; i < len(nums)-k+1; i++ {
		if nums[i]-nums[i+k-1] < min {
			min = nums[i] - nums[i+k-1]
		}
	}
	return min
}
