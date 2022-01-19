package main

import "fmt"

func main() {
	fmt.Println(containsNearbyDuplicate([]int{1, 2, 3, 1}, 3))       // true
	fmt.Println(containsNearbyDuplicate([]int{1, 0, 1, 1}, 1))       // true
	fmt.Println(containsNearbyDuplicate([]int{1, 2, 3, 1, 2, 3}, 2)) //false
}

func containsNearbyDuplicate(nums []int, k int) bool {
	hashMap := map[int]int{}
	for i, v := range nums {
		if val, ok := hashMap[v]; ok && i-val <= k {
			return true
		}
		hashMap[v] = i
	}
	return false
}
