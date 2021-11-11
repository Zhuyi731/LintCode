package main

import "fmt"

func main() {
	fmt.Println(kInversePairs(3, 0)) // 1
	fmt.Println(kInversePairs(3, 1)) // 2
}

var dp [][]int

func init() {
	mod := int(1e9 + 7)
	dp = make([][]int, 1001)
	dp[0] = make([]int, 1001)
	dp[0][0] = 1
	for i := 1; i <= 1000; i++ {
		dp[i] = make([]int, 1001)
		dp[i][0] = 1
		for j := 1; j <= 1000; j++ {
			dp[i][j] = (dp[i][j-1] + dp[i-1][j])
			if j >= i {
				dp[i][j] = (dp[i][j] - dp[i-1][j-i])
			}
			if dp[i][j] >= mod {
				dp[i][j] -= mod
			} else if dp[i][j] < 0 {
				dp[i][j] += mod
			}
		}
	}
}

func kInversePairs(n int, k int) int {
	return dp[n][k]
}
