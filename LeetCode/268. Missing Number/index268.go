package main

import "fmt"

func main() {
	fmt.Println(missingNumber([]int{3, 0, 1}))
	fmt.Println(missingNumber([]int{0, 1}))
	fmt.Println(missingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}))
	fmt.Println(missingNumber([]int{0}))
}

func missingNumber(nums []int) int {
	l := len(nums)
	i := 0
	for {
		if i == l {
			for k, v := range nums {
				if v == -1 {
					return k
				}
			}
			return l
		} else if nums[i] == i || nums[i] == -1 {
			i++
		} else if nums[i] == l {
			nums = append(nums, nums[i])
			nums[i] = -1
			i++
		} else {
			nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
		}
	}
}
