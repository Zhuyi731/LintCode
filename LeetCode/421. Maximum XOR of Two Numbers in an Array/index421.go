package main

import "fmt"

func main() {
	fmt.Println(findMaximumXOR([]int{5, 25, 3, 10, 2, 8}))                             // 28
	fmt.Println(findMaximumXOR([]int{0}))                                              // 0
	fmt.Println(findMaximumXOR([]int{2, 4}))                                           //6
	fmt.Println(findMaximumXOR([]int{8, 10, 2}))                                       // 10
	fmt.Println(findMaximumXOR([]int{14, 70, 53, 83, 49, 91, 36, 80, 92, 51, 66, 70})) // 127
}

type TrieTree struct {
	Left  *TrieTree
	Right *TrieTree
}

const maxBit = 30

func (t *TrieTree) Add(num int) {
	cur := t
	for i := maxBit; i >= 0; i-- {
		bit := num >> i & 1
		if bit == 1 {
			if cur.Left == nil {
				cur.Left = &TrieTree{}
			}
			cur = cur.Left
		} else {
			if cur.Right == nil {
				cur.Right = &TrieTree{}
			}
			cur = cur.Right
		}
	}
}

func (t *TrieTree) GetVal(num int) int {
	cur := t
	result := 0
	for i := maxBit; i >= 0; i-- {
		bit := num >> i

		if bit%2 == 1 {
			if cur.Right == nil {
				cur = cur.Left
			} else {
				cur = cur.Right
				result += 1 << i
			}
		} else {
			if cur.Left == nil {
				cur = cur.Right
			} else {
				cur = cur.Left
				result += 1 << i
			}
		}
	}
	return result
}

/**
 * @param {number[]} nums
 * @return {number}
 */
func findMaximumXOR(nums []int) int {
	t := TrieTree{}
	for _, v := range nums {
		t.Add(v)
	}
	// t.print()
	maxVal := 0
	for _, v := range nums {
		tmp := t.GetVal(v)
		if tmp > maxVal {
			maxVal = tmp
		}
	}
	return maxVal
}
