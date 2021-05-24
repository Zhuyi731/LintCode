package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maximizeXor([]int{5, 8, 0, 3, 2, 10, 9, 2, 4, 5}, [][]int{{3, 8}}))
	fmt.Println(maximizeXor([]int{0, 1, 2, 3, 4}, [][]int{{3, 1}, {1, 3}, {5, 6}}))
	fmt.Println(maximizeXor([]int{5, 2, 4, 6, 6, 3}, [][]int{{12, 4}, {8, 1}, {6, 3}}))
}

type BinaryTrieTree struct {
	Left   *BinaryTrieTree
	Right  *BinaryTrieTree
	MinVal int
}

const max = 4

func (b *BinaryTrieTree) Insert(val int) {
	bt := b
	for i := max; i >= 0; i-- {
		bit := val >> i & 1
		if val < bt.MinVal {
			bt.MinVal = val
		}
		if bit == 0 {
			if bt.Left == nil {
				bt.Left = &BinaryTrieTree{
					MinVal: val,
				}
			}
			bt = bt.Left
		} else {
			if bt.Right == nil {
				bt.Right = &BinaryTrieTree{
					MinVal: val,
				}
			}
			bt = bt.Right
		}
	}
}

func (b *BinaryTrieTree) Query(val int, limit int) int {
	bt := b
	result := 0
	for i := max; i >= 0; i-- {
		if bt.Left == nil && bt.Right == nil {
			return result
		}
		if bt.MinVal > limit {
			return -1
		}
		bit := val >> i & 1
		if bit == 1 {
			// 往左
			if bt.Left != nil && bt.Left.MinVal <= limit {
				bt = bt.Left
				result = result*2 + 1
			} else if bt.Right != nil && bt.Right.MinVal <= limit {
				bt = bt.Right
				result = result*2 + 0
			} else {
				return -1
			}
		} else {
			if bt.Right != nil && bt.Right.MinVal <= limit {
				bt = bt.Right
				result = result*2 + 1
			} else if bt.Left != nil && bt.Left.MinVal <= limit {
				bt = bt.Left
				result = result*2 + 0
			} else {
				return -1
			}
		}
	}
	return result
}

func maximizeXor(nums []int, queries [][]int) []int {
	bt := BinaryTrieTree{
		MinVal: math.MaxInt32,
	}
	for _, v := range nums {
		bt.Insert(v)
	}
	result := make([]int, len(queries))
	for i, q := range queries {
		tmp := bt.Query(q[0], q[1])
		result[i] = tmp
	}

	return result
}
