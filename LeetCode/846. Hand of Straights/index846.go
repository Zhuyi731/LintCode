package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(isNStraightHand([]int{1, 2, 3, 6, 2, 3, 4, 7, 8}, 3))
	fmt.Println(isNStraightHand([]int{1, 2, 3, 5, 2, 3, 4, 7, 8}, 3))
	fmt.Println(isNStraightHand([]int{4, 2, 3, 6, 2, 3, 4, 7, 8}, 3))
	fmt.Println(isNStraightHand([]int{9, 10, 3, 6, 2, 11, 4, 7, 8}, 3))
	fmt.Println(isNStraightHand([]int{1, 2, 3, 4, 5}, 4))
}

type record struct {
	ct  int
	num int
}

func isNStraightHand(hand []int, groupSize int) bool {
	if len(hand)%groupSize != 0 {
		return false
	}
	if groupSize == 1 {
		return true
	}

	sort.Ints(hand)
	// 然后收集数量
	nums := []record{}
	currentNum := hand[0]
	currentCt := 0
	for _, v := range hand {
		if v == currentNum {
			currentCt++
		} else {
			nums = append(nums, record{ct: currentCt, num: currentNum})
			currentCt = 1
			currentNum = v
		}
	}
	nums = append(nums, record{ct: currentCt, num: currentNum})
	index := 0
	for {
		if index >= len(nums) {
			break
		}
		firstCt := nums[index].ct
		for i := index + 1; i < index+groupSize; i++ {
			if i >= len(nums) {
				return false
			}
			if nums[i].num != nums[i-1].num+1 {
				return false
			}
			if nums[i].ct < firstCt {
				return false
			}
			nums[i].ct -= firstCt
		}
		nums[index].ct = 0
		for {
			if index >= len(nums) {
				return true
			}
			if nums[index].ct == 0 {
				index++
			} else {
				break
			}
		}
	}
	return true
}
