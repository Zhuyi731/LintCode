package main

import "fmt"

func main() {
	fmt.Println(minDistance("horse", "ros"))
	fmt.Println(minDistance("intention", "execution"))
}

const INT_MAX = int(^uint(0) >> 1)

func minDistance(word1 string, word2 string) int {
	len1 := len(word1)
	len2 := len(word2)
	var dp [][]int
	dp = make([][]int, len1+1)

	//c初始化
	for i := 0; i <= len1; i++ {
		dp[i] = make([]int, len2+1)
		dp[i][0] = i
	}
	for i := 0; i <= len2; i++ {
		dp[0][i] = i
	}

	for i := 1; i <= len1; i++ {
		for j := 1; j <= len2; j++ {
			min := dp[i-1][j] + 1
			min1 := dp[i][j-1] + 1
			if min1 < min {
				min = min1
			}
			var min2 int
			if word1[i-1] == word2[j-1] {
				min2 = dp[i-1][j-1]
			} else {
				min2 = dp[i-1][j-1] + 1
			}
			if min2 < min {
				min = min2
			}
			dp[i][j] = min
		}
	}
	return dp[len1][len2]
}
