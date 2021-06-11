package main

import "fmt"

func main() {
	fmt.Println(change(5, []int{1, 2, 5}))
	fmt.Println(change(3, []int{2}))
	fmt.Println(change(3, []int{3}))
	fmt.Println(change(26, []int{1, 2, 3, 6}))
}

func change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for _, v := range coins {
		for i := 1; i <= amount; i++ {
			if i-v >= 0 {
				dp[i] += dp[i-v]
			}
		}

	}

	return dp[amount]
}
