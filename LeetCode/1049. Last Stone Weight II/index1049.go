package main

import "fmt"

func main() {
	fmt.Println(lastStoneWeightII([]int{2, 7, 4, 1, 8, 1}))
	fmt.Println(lastStoneWeightII([]int{31, 26, 33, 21, 40}))
	fmt.Println(lastStoneWeightII([]int{1, 2}))
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func lastStoneWeightII(stones []int) int {
	dp := make([][]bool, len(stones)+1)
	dp[0] = []bool{true}
	sum := 0
	for i, v := range stones {
		sum += v
		dp[i+1] = make([]bool, sum+1)
		for j := 0; j <= sum; j++ {
			m := abs(j, v)
			if m < len(dp[i]) && dp[i][m] {
				dp[i+1][j] = true
			}
			if len(dp[i]) > j+v && dp[i][j+v] {
				dp[i+1][j] = true
			}
		}
	}

	for i, v := range dp[len(stones)] {
		if v {
			return i
		}
	}
	return -1
}
