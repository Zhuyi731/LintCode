package main

import "fmt"

func main() {
	fmt.Println(countSubstrings("aaa"))
	fmt.Println(countSubstrings("abc"))
}

func countSubstrings(s string) int {
	lens := len(s)
	dp := make([]int, lens)
	for index := range s {
		// 奇数循环
		ct := 1
		dp[index] = 0
		for {
			if index-ct >= 0 && index+ct < lens && s[index-ct] == s[index+ct] {
				ct++
			} else {
				dp[index] = ct
				break
			}
		}

		// 偶数循环
		if index < lens-1 && s[index] == s[index+1] {
			ct = 1
			for {
				if index-ct >= 0 && index+1+ct < lens && s[index-ct] == s[index+1+ct] {
					ct++
				} else {
					dp[index] += ct
					break
				}
			}
		}
	}

	result := 0
	for _, v := range dp {
		result += v
	}
	return result
}
