package main

import "fmt"

func main() {
	fmt.Println(longestSubsequence([]int{1, 2, 3, 4}, 1))
	fmt.Println(longestSubsequence([]int{1, 3, 5, 7}, 1))
	fmt.Println(longestSubsequence([]int{1, 5, 7, 8, 5, 3, 4, 2, 1}, -2))
}

func longestSubsequence(arr []int, difference int) (m int) {
	dp := map[int]int{}
	for _, v := range arr {
		dp[v] = dp[v-difference] + 1
		if dp[v] > m {
			m = dp[v]
		}
	}
	return m
}
