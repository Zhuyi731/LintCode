package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findRadius([]int{1, 2, 3}, []int{2}))            //2
	fmt.Println(findRadius([]int{1, 2, 3, 5, 15}, []int{2, 30})) //13
	fmt.Println(findRadius([]int{4, 1, 2, 3}, []int{1, 4}))      //2
	fmt.Println(findRadius([]int{1, 5}, []int{2}))
}
func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func findRadius(houses []int, heaters []int) int {
	maxDist := 0
	sort.Ints(houses)
	sort.Ints(heaters)
	curIndex := 0

	var cmp func(i, v int) int
	cmp = func(i, v int) int {
		if i == 0 {
			return abs(v - heaters[0])
		}
		if abs(v-heaters[i]) > abs(v-heaters[i-1]) {
			return abs(v - heaters[i-1])
		}
		return abs(v - heaters[i])
	}

	for _, v := range houses {
		if v < heaters[curIndex] {
			dist := cmp(curIndex, v)
			if maxDist < dist {
				maxDist = dist
			}
		} else if v > heaters[curIndex] {
			for {
				if curIndex >= len(heaters)-1 || v < heaters[curIndex] {
					break
				}
				curIndex++
			}
			dist := cmp(curIndex, v)
			if maxDist < dist {
				maxDist = dist
			}
		}
	}
	return maxDist
}
