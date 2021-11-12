package main

import "fmt"

func main() {
	fmt.Println(getMoneyAmount(1))   //0
	fmt.Println(getMoneyAmount(2))   //1
	fmt.Println(getMoneyAmount(3))   //2
	fmt.Println(getMoneyAmount(4))   //4
	fmt.Println(getMoneyAmount(10))  //16
	fmt.Println(getMoneyAmount(40))  //119
	fmt.Println(getMoneyAmount(60))  //214
	fmt.Println(getMoneyAmount(80))  //295
	fmt.Println(getMoneyAmount(199)) //946
}

var cost [][]int

func init() {
	cost = make([][]int, 202)
	for i := 0; i <= 201; i++ {
		cost[i] = make([]int, 202)
	}
	for i := 1; i <= 200; i++ {
		for j := 1; j <= 200; j++ {
			if i+j > 200 {
				break
			}
			cost[j][j+i] = 0xfffffff
			for k := j; k <= i+j; k++ {
				c := max(cost[j][k-1], cost[k+1][j+i]) + k
				if c < cost[j][i+j] {
					cost[j][j+i] = c
				}
			}
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getMoneyAmount(n int) int {
	return cost[1][n]
}
