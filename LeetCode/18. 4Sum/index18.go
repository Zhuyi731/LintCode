package main 

import (
	"fmt"
	"sort"
)

func main (){
	fmt.Println(fourSum([]int{1, 0,0,0, -1, 0, -2, 2},0))
	fmt.Println(fourSum([]int{1,0,-1,3,0,1,0,-2,2},1))
}

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	var ans [][]int
	numsLen := len(nums)

	for i := 0;i<numsLen-3;i++ {
		for j:=i+1;j<numsLen-2;j++{
			for k:=j+1;k<numsLen-1;k++{
				for t := k+1 ;t<numsLen;t++{
					cur := nums[i] + nums[j] + nums[k]+nums[t]
					if cur == target {
						temp := append([]int{},nums[i],nums[j],nums[k],nums[t])
						ans = append(ans,temp)
					}
				}
			}
		}
	}
	

	ansLen := len(ans)
	for i:=0;i<ansLen-1;i++{
		for j:=i+1;j<ansLen;j++{
			if ans[i][0] == ans[j][0] &&
			ans[i][1] == ans[j][1] &&
			ans[i][2] == ans[j][2] &&
			ans[i][3] == ans[j][3] {
				ans = append(ans[:i],ans[i+1:]...)
				i--
				ansLen--
				break;
			}
		}
	}

	return ans
}

func abs(val int) int{
	if val < 0{
		val = -val
	}
	return val
}
// Given an array nums of n integers and an integer target, are there elements a, b, c, and d in nums such that a + b + c + d = target? Find all unique quadruplets in the array which gives the sum of target.

// Note:

// The solution set must not contain duplicate quadruplets.

// Example:

// Given array nums = [1, 0, -1, 0, -2, 2], and target = 0.

// A solution set is:
// [
//   [-1,  0, 0, 1],
//   [-2, -1, 1, 2],
//   [-2,  0, 0, 2]
// ]