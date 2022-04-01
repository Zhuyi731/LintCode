package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(canReorderDoubled([]int{-2, -6, -3, 4, -4, 2}))
	fmt.Println(canReorderDoubled([]int{2, 4, 0, 0, 8, 1}))
	fmt.Println(canReorderDoubled([]int{4, -2, 2, -4}))
	fmt.Println(canReorderDoubled([]int{3, 1, 3, 6}))
	fmt.Println(canReorderDoubled([]int{2, 1, 3, 6}))
	fmt.Println(canReorderDoubled([]int{2, 1, 2, 6}))
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
func canReorderDoubled(arr []int) bool {
	mp := map[int]int{}
	for _, v := range arr {
		mp[v]++
	}
	if mp[0]%2 != 0 {
		return false
	}
	delete(mp, 0)

	sort.Slice(arr, func(i, j int) bool {
		if arr[i]*arr[j] < 0 {
			return arr[i] < arr[j]
		} else {
			return abs(arr[i]) < abs(arr[j])
		}
	})
	for i := 0; i < len(arr); i++ {
		key := arr[i]
		if mp[key] == 0 {
			continue
		}
		if mp[key] > mp[key*2] {
			return false
		} else {
			mp[key*2] -= mp[key]
			mp[key] = 0
		}
	}
	return true
}
