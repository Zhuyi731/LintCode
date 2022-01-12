package main

import "fmt"

func main() {
	fmt.Println(increasingTriplet([]int{5, 1, 6}))
	fmt.Println(increasingTriplet([]int{1, 2, 3, 4, 5}))
	fmt.Println(increasingTriplet([]int{5, 4, 3, 2, 1}))
	fmt.Println(increasingTriplet([]int{2, 1, 5, 0, 4, 6}))
}

func increasingTriplet(nums []int) bool {
	if len(nums) < 3 {
		return false
	}
	first := nums[0]
	second := 1<<32 - 1
	for i, v := range nums {
		if i == 0 {
			continue
		}
		if v > second {
			return true
		}
		if v < second && v > first {
			second = v
		}
		if v < first {
			first = v
		}
	}
	return false
}
