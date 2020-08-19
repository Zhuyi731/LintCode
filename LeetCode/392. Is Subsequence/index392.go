package main

import "fmt"

func main() {
	fmt.Println(isSubsequence("aaaaaa", "aaa"))
	fmt.Println(isSubsequence("abc", "ahbgdc"))
	fmt.Println(isSubsequence("axc", "ahbgdc"))
	fmt.Println(isSubsequence("", ""))
	fmt.Println(isSubsequence("a", ""))
	fmt.Println(isSubsequence("", "c"))
}

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	if len(t) == 0 {
		return false
	}
	dp := make([][]bool, len(s))

	for i, valS := range s {
		dp[i] = make([]bool, len(t))

		for j, valT := range t {
			if valS == valT {
				if i == 0 || j == 0 {
					if i == 0 {
						dp[i][j] = true
					} else {
						dp[i][j] = false
					}
				} else {
					dp[i][j] = dp[i-1][j-1] || dp[i][j-1]
				}
			} else {
				if j == 0 {
					dp[i][j] = false
				} else {
					dp[i][j] = dp[i][j-1] || false
				}
			}
		}
	}
	return dp[len(s)-1][len(t)-1]
}
