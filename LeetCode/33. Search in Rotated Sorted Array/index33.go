package main 

import (
	"fmt"
	"regexp"
)

// Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.

// (i.e., [0,1,2,4,5,6,7] might become [4,5,6,7,0,1,2]).

// You are given a target value to search. If found in the array return its index, otherwise return -1.

// You may assume no duplicate exists in the array.

// Your algorithm's runtime complexity must be in the order of O(log n).

// Example 1:

// Input: nums = [4,5,6,7,0,1,2], target = 0
// Output: 4
// Example 2:

// Input: nums = [4,5,6,7,0,1,2], target = 3
// Output: -1
func main(){
	match,_ := regexp.MatchString("models/.*","models/43b2b961786cdf294254088b3850300d.zip")
	match2,_ := regexp.MatchString("models/.*","pics/83a168c7ec4e74c1f8f8fe01a84c5ef2.jpg")
	fmt.Println(match,match2)
	r,_ := regexp.Compile("models/([0-9a-zA-Z]*)\\..*")
	fmt.Println(r.FindStringSubmatch("pics/83a168c7ec4e74c1f8f8fe01a84c5ef2.jpg"))
	fmt.Println(r.FindStringSubmatch("models/43b2b961786cdf294254088b3850300d.zip"))
	
	// fmt.Println(search([]int{4,5,6,7,0,1,2},0))
	// fmt.Println(search([]int{4,5,6,7,0,1,2},3))
}

// func search(nums []int, target int) int {
// 	var l ,r ,mid,index int 
	
// 	for {
// 		if l<r{
			
// 		}
// 	}

// 	return index
// }