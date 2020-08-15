package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(jump([]int{2, 3, 1, 1, 4}))
}

func jump(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = 0
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < nums[i]+i+1 && j < len(nums); j++ {
			if dp[j] != 0 {
				dp[j] = int(math.Min(float64(dp[j]), float64(dp[i]+1)))
			} else {
				dp[j] = dp[i] + 1
			}
		}
	}
	return dp[len(nums)-1]
}
