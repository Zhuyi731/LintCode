package main

import "fmt"

func main() {
	fmt.Println(numFriendRequests([]int{16, 16}))                //2
	fmt.Println(numFriendRequests([]int{16, 17, 18}))            //2
	fmt.Println(numFriendRequests([]int{20, 30, 100, 110, 120})) // 3
}

// x> y > x/2 +7
func numFriendRequests(ages []int) int {
	cnt, sum := [121]int{}, [121]int{}
	for _, v := range ages {
		cnt[v]++
	}
	for i := 1; i < 121; i++ {
		sum[i] += sum[i-1] + cnt[i]
	}
	result := 0
	for i := 15; i < 121; i++ {
		result += cnt[i] * (sum[i] - sum[(i/2+7)] - 1)
	}
	return result
}
