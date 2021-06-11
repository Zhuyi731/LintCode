package main

import "fmt"

func main() {
	fmt.Println(numSquares(12))
	fmt.Println(numSquares(13))
	fmt.Println(numSquares(1322))
}

var nums []int

func init() {
	nums = make([]int, 0)
	i := 1
	for {
		s := i * i
		if s >= 10000 {
			break
		} else {
			nums = append(nums, s)
		}
		i++
	}
}

func numSquares(n int) int {
	dp := make([]int, n+1)
	dp[0] = 0
	for j := 1; j <= n; j++ {
		dp[j] = 100000
	}
	for i := len(nums) - 1; i >= 0; i-- {
		for j := 0; j <= n; j++ {
			if j-nums[i] >= 0 {
				if dp[j] > dp[j-nums[i]]+1 {
					dp[j] = dp[j-nums[i]] + 1
				}
			}
		}
	}
	return dp[n]
}
