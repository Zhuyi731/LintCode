package main

import "fmt"

func main() {
	removeDuplicates([]int{1, 1, 1, 2, 2, 3})
	removeDuplicates([]int{0, 0, 1, 1, 1, 1, 2, 3, 3})
}

func removeDuplicates(nums []int) int {
	l := len(nums)

	for i := 1; i < l; i++ {
		if nums[i] == nums[i-1] && i+1 < l && nums[i+1] == nums[i] {
			// 删除当前数
			for j := i; j < l-1; j++ {
				nums[j] = nums[j+1]
			}
			l--
			i--
		}
	}

	fmt.Println(l)
	return l
}
