package main

import "fmt"

func main() {
	fmt.Println(countQuadruplets([]int{1, 2, 3, 6}))
	fmt.Println(countQuadruplets([]int{3, 3, 6, 4, 5}))
	fmt.Println(countQuadruplets([]int{1, 1, 1, 3, 5}))
}

func countQuadruplets(nums []int) int {
	countMap := map[int][]int{}
	for i, v := range nums {
		if _, ok := countMap[v]; ok {
			countMap[v] = append(countMap[v], i)
		} else {
			countMap[v] = []int{i}
		}
	}

	ct := 0
	l := len(nums)
	for a := 0; a < l-2; a++ {
		for b := a + 1; b < l-1; b++ {
			for c := b + 1; c < l; c++ {
				currentSum := nums[a] + nums[b] + nums[c]
				arr := countMap[currentSum]
				for i, v := range arr {
					if v > c {
						ct += len(arr) - i
						break
					}
				}
			}
		}
	}
	return ct
}
