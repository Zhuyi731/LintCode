package main

import (
	"fmt"
)

func main() {
	fmt.Println(totalHammingDistance([]int{4, 14, 2}))
	fmt.Println(totalHammingDistance([]int{4, 14, 4}))
	fmt.Println(totalHammingDistance([]int{4, 14, 4, 3, 123, 6, 78}))
}

func totalHammingDistance(nums []int) int {
	resultMap := map[int]int{}
	result := 0
	for _, v := range nums {
		ct := 0
		for v>>ct > 0 {
			if v>>ct&1 == 1 {
				resultMap[ct]++
			}
			ct++
		}
	}
	ln := len(nums)
	for _, v := range resultMap {
		result += v * (ln - v)
	}
	return result
}
