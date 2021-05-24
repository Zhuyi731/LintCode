package main

import "fmt"

func main() {
	fmt.Println(strangePrinter("cabac"))
	fmt.Println(strangePrinter("aaabbb"))
	fmt.Println(strangePrinter("aba"))
}

func strangePrinter(s string) int {
	lens := len(s)
	dp := make([][]int, lens)
	for i := 0; i < lens; i++ {
		dp[i] = make([]int, lens)
		dp[i][i] = 1
	}

	for i := lens - 1; i >= 0; i-- {
		for j := i + 1; j < lens; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i][j-1]
			} else {
				max := 0xfffffff
				for k := i; k < j; k++ {
					cur := dp[i][k] + dp[k+1][j]
					if cur < max {
						max = cur
					}
				}
				dp[i][j] = max
			}
		}
	}

	return dp[0][lens-1]
}
