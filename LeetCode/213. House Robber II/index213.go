package main

import "fmt"

func main() {
	rob([]int{1, 2})
	rob([]int{2, 1})
	rob([]int{2, 3, 2})
	rob([]int{1, 2, 3, 1})
}

func max(nums ...int) int {
	max := 0
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	return max
}

func rob(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}
	if l == 1 {
		return nums[0]
	}
	dp := make([][][]int, l)
	for i := 0; i < l; i++ {
		dp[i] = make([][]int, l)
		for j := 0; j < l; j++ {
			dp[i][j] = make([]int, 2)
		}
	}

	// dp[i][j][k] 表示  到第i间房 j表示是否抢劫该刚 k表示是否抢劫第一间房的最大值

	dp[0][0][0] = 0
	dp[0][0][1] = 0
	dp[0][1][0] = 0
	dp[0][1][1] = nums[0]
	for i := 1; i < l; i++ {
		if i > 1 {
			dp[i][0][0] = max(dp[i-1][0][0], dp[i-1][1][0])
			dp[i][0][1] = max(dp[i-1][1][1], dp[i-1][0][1])
			dp[i][1][0] = max(dp[i-2][1][0]+nums[i], dp[i-1][0][0]+nums[i])
			dp[i][1][1] = max(dp[i-2][1][1]+nums[i], dp[i-1][0][1]+nums[i])
		} else {
			dp[i][0][0] = 0
			dp[i][0][1] = nums[0]
			dp[i][1][0] = nums[1]
			dp[i][1][1] = 0
		}
	}
	fmt.Println(max(dp[l-1][1][0], dp[l-1][0][0], dp[l-1][0][1]))
	return max(dp[l-1][1][0], dp[l-1][0][0], dp[l-1][0][1])
}
