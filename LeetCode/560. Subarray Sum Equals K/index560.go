package main

import "fmt"

func main() {
	fmt.Println(subarraySum([]int{1}, 0))
	fmt.Println(subarraySum([]int{1, 1, 1}, 2))
	fmt.Println(subarraySum([]int{1, 2, 3}, 3))
}

func subarraySum(nums []int, k int) int {
	cur, result := 0, 0
	mapper := map[int]int{
		0: 1,
	}
	for _, v := range nums {
		cur = cur + v
		result += mapper[cur-k]
		mapper[cur]++
	}
	return result
}
