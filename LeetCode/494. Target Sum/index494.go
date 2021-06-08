package main

import "fmt"

func main() {
	fmt.Println(findTargetSumWays([]int{1, 1, 1, 1, 1}, 3))
	fmt.Println(findTargetSumWays([]int{1}, 1))
}

func findTargetSumWays(nums []int, target int) int {
	dp := make([]map[int]int, len(nums)+1)
	dp[0] = map[int]int{0: 1}

	for i := 0; i < len(nums); i++ {
		dp[i+1] = map[int]int{}
		for j := -1000; j <= 1000; j++ {
			dp[i+1][j] = dp[i][j-nums[i]] + dp[i][j+nums[i]]
		}
	}
	return dp[len(nums)][target]
}
