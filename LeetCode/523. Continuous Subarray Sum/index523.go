package main

import (
	"fmt"
)

func main() {
	fmt.Println(checkSubarraySum([]int{0}, 1))
	fmt.Println(checkSubarraySum([]int{23, 2, 4, 6, 7}, 6))
	fmt.Println(checkSubarraySum([]int{23, 2, 6, 4, 7}, 6))
	fmt.Println(checkSubarraySum([]int{23, 2, 6, 4, 7}, 13))
}
func checkSubarraySum(nums []int, k int) bool {
	sum := make([]int, len(nums)+1)
	hashMap := map[int]int{
		0: -1,
	}
	for i, v := range nums {
		sum[i+1] = (sum[i] + v) % k
		if index, ok := hashMap[sum[i+1]]; ok {
			if i-index > 1 {
				return true
			}
		} else {
			hashMap[sum[i+1]] = i
		}
	}

	return false
}
