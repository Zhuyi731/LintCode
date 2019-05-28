package main 

import (
	"fmt"
	"sort"
)

func main (){
var a []int

	a = []int{1,3,2}
	nextPermutation(a)
	fmt.Println(a)
	 a = []int{2,3,1}
	nextPermutation(a)
	fmt.Println(a)
	a = []int{3,2,1}
	nextPermutation(a)
	fmt.Println(a)
	a = []int{1,1,5}
	nextPermutation(a)
	fmt.Println(a)
}

func nextPermutation(nums []int)  {
    numsLen := len(nums)
	if numsLen == 1{
		return
	}
	if numsLen == 2{
		nums[0],nums[1]=nums[1],nums[0]
		return
	}
	
	var finded bool = false
	for i:=numsLen-1;i>0;i--{
		if nums[i-1] < nums[i]{
			//就是你了
			for j:= numsLen-1 ;j >= i ;j--{
				if nums[j] > nums[i-1]{
					//说明  i-1  ,j 的位置需要调换   
					nums[i-1],nums[j] = nums[j],nums[i-1]
					finded = true
					//然后对i ~ 末尾 排序
					sort.Ints(nums[i:])
					break
				}
			}
			break;
		}
	}
	if !finded {
		//如果是最大值，则变成最小值
		sort.Ints(nums)
	}
}