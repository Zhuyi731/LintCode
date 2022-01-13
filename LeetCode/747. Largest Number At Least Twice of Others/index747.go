package main

import "fmt"

func main() {
	fmt.Println(dominantIndex([]int{0, 0, 2, 2})) // 0
	fmt.Println(dominantIndex([]int{3, 6, 1, 0})) //1
	fmt.Println(dominantIndex([]int{1, 2, 3, 4})) //-1
	fmt.Println(dominantIndex([]int{1}))          // 0
	fmt.Println(dominantIndex([]int{}))           // 0
}

func dominantIndex(nums []int) int {
	first, second := -(1<<31 - 1), -(1<<31 - 1)
	result := -1
	for i, v := range nums {
		if v > first {
			second = first
			first = v
			result = i
		} else if v > second && v != first {
			second = v
		}
	}
	if second <= 0 || first/second >= 2 {
		return result
	}
	return -1
}
