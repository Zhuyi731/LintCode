package main

import "fmt"

func main() {
	fmt.Println(majorityElement([]int{3, 2, 3}))
	fmt.Println(majorityElement([]int{1, 1, 1, 3, 3, 2, 2, 2}))
	fmt.Println(majorityElement([]int{1}))
	fmt.Println(majorityElement([]int{1, 2, 3}))
	fmt.Println(majorityElement([]int{1, 1, 1, 3, 3, 3, 2, 2, 2}))
}

func majorityElement(nums []int) []int {
	l := len(nums)/3 + 1
	result := []int{}
	hMap := map[int]int{}
	for _, v := range nums {
		hMap[v]++
		if hMap[v] == l {
			result = append(result, v)
		}
	}
	return result
}
