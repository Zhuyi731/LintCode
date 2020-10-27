package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(smallerNumbersThanCurrent([]int{8, 1, 2, 2, 3}))
	fmt.Println(smallerNumbersThanCurrent([]int{6, 5, 4, 8}))
	fmt.Println(smallerNumbersThanCurrent([]int{7, 7, 7, 7}))
}

type Item struct {
	Val   int
	Index int
}

type NewNums []Item

func (n NewNums) Len() int {
	return len(n)
}

func (n NewNums) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

func (n NewNums) Less(i, j int) bool {
	return n[i].Val < n[j].Val
}

func smallerNumbersThanCurrent(nums []int) []int {
	n := make([]Item, len(nums))
	for i, v := range nums {
		n[i] = Item{
			Val:   v,
			Index: i,
		}
	}

	sort.Sort(NewNums(n))
	result := make([]int, len(nums))

	for i, v := range n {
		if i > 0 && v.Val == n[i-1].Val {
			result[v.Index] = result[n[i-1].Index]
		} else {
			result[v.Index] = i
		}
	}

	return result
}
