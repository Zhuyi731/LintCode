package main

import (
	"fmt"
)

func main() {
	sortColors([]int{2})
	sortColors([]int{2, 1, 1, 2, 2})
	sortColors([]int{2, 0, 0, 0, 2})
	sortColors([]int{2, 0, 2, 1, 1, 0})
	sortColors([]int{2, 0, 1})
}

func sortColors(nums []int) {
	hsMap := map[int]int{
		0: 0,
		1: 0,
		2: 0,
	}

	for _, v := range nums {
		hsMap[v]++
	}

	cur := 0
	total := hsMap[cur]
	if total == 0 {
		cur++
		total = hsMap[cur]
		if hsMap[cur] == 0 {
			cur++
			total = hsMap[cur]
		}
	}
	for i := 0; i < len(nums); i++ {
		if i >= total {
			cur++
			if hsMap[cur] == 0 {
				cur++
			}
			total += hsMap[cur]
		}
		nums[i] = cur
	}

	fmt.Println(nums)
}
