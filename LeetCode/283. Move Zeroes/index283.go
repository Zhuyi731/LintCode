package main

import "fmt"

func main() {
	moveZeroes([]int{1, 2, 0})
	moveZeroes([]int{1, 2})
	moveZeroes([]int{0, 1, 0, 3, 12})
	moveZeroes([]int{0, 0})
}

func moveZeroes(nums []int) {
	if len(nums) < 2 {
		return
	}
	posPtr := -1

	for i, v := range nums {
		if v == 0 {
			if i > posPtr {
				posPtr = i
			}
			j := posPtr + 1
			for {
				if j == len(nums) {
					return
				}
				if nums[j] != 0 {
					posPtr = j
					nums[i], nums[j] = nums[j], nums[i]
					break
				}
				j++
			}
		}
	}
	fmt.Println(nums)
}
