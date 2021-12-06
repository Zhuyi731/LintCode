package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Println(findNthDigit(11))  //0
	fmt.Println(findNthDigit(100)) //5
	fmt.Println(findNthDigit(3))   //3
	fmt.Println(findNthDigit(9))   //9
	fmt.Println(findNthDigit(10))  //1
	fmt.Println(findNthDigit(134)) //7
}

var nums = []int{}

func init() {
	pre := 1
	for i := 1; i < 12; i++ {
		t := int(math.Pow10(i))
		nums = append(nums, t-pre)
		pre = t
	}
}

func findNthDigit(n int) int {
	i := 0
	sum := 0
	for ; i < 12; i++ {
		if n > nums[i]*(i+1) {
			n = n - nums[i]*(i+1)
			sum += nums[i]
		} else {
			break
		}
	}

	c, r := (n-1)/(i+1)+1, (n-1)%(i+1)+1
	targetNumber := sum + c
	//获取targetNumber的第r位数即可
	targetNumberStr := strconv.Itoa(targetNumber)
	result, _ := strconv.Atoi(string(targetNumberStr[r-1]))
	return result

}
