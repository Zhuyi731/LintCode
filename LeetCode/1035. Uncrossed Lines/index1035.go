package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxUncrossedLines([]int{1, 4, 2}, []int{1, 2, 4}))
	fmt.Println(maxUncrossedLines([]int{2, 5, 1, 2, 5}, []int{10, 5, 2, 1, 5, 2}))
	fmt.Println(maxUncrossedLines([]int{1, 3, 7, 1, 7, 5}, []int{1, 9, 2, 5, 1}))
}

const mixVal = math.MinInt64

func max(val ...int) int {
	maxVal := mixVal
	for _, v := range val {
		if maxVal < v {
			maxVal = v
		}
	}
	return maxVal
}
func btoi(i bool) int {
	if i {
		return 1
	}
	return 0
}

func maxUncrossedLines(nums1 []int, nums2 []int) int {
	dp := make([][]int, len(nums1))
	for i, v1 := range nums1 {
		dp[i] = make([]int, len(nums2))
		for j, v2 := range nums2 {
			if i == 0 {
				if j == 0 {
					dp[i][j] = btoi(v1 == v2)
				} else {
					dp[i][j] = max(btoi(v1 == v2), dp[i][j-1])
				}
			} else if j == 0 {
				dp[i][j] = max(btoi(v1 == v2), dp[i-1][j])
			} else {
				dp[i][j] = max(dp[i-1][j-1]+btoi(v1 == v2), dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[len(nums1)-1][len(nums2)-1]
}
