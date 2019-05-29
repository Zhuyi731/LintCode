package main 

import (
	"fmt"
	"sort"
)

func main (){
	fmt.Println(threeSumClosest([]int{-1, 2, 1, -4},1))
	fmt.Println(threeSumClosest([]int{-1, 0, 2, 1, -4},1))
}

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	var cur,min,ans int
	min = 99999999
	var numsLen = len(nums)
	for i := 0;i<numsLen-2;i++ {
		for j:=i+1;j<numsLen-1;j++{
			for k:=j+1;k<numsLen;k++{
				cur = nums[i] + nums[j] + nums[k]
				if min !=-1 && abs(cur-target)<min{
					min=abs(cur-target)
					ans=cur
				}
			}
		}
	}

    return ans
}

func abs(val int) int{
	if val < 0{
		val = 0-val
	}
	return val
}

//该题可以进行剪枝优化   这种解法较慢
/***
Given an array nums of n integers and an integer target, find three integers in nums such that the sum is closest to target. Return the sum of the three integers. You may assume that each input would have exactly one solution.

Example:

Given array nums = [-1, 2, 1, -4], and target = 1.

The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).
**/