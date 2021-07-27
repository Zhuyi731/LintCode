package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

const NULL = 0X7FFFFFFF
const null = 0x7FFFFFFF

func Ints2TreeNode(ints []int) *TreeNode {
	n := len(ints)
	if n == 0 {
		return nil
	}

	root := &TreeNode{
		Val: ints[0],
	}

	queue := make([]*TreeNode, 1, n*2)
	queue[0] = root

	i := 1
	for i < n {
		node := queue[0]
		queue = queue[1:]

		if i < n && ints[i] != NULL {
			node.Left = &TreeNode{Val: ints[i]}
			queue = append(queue, node.Left)
		}
		i++

		if i < n && ints[i] != NULL {
			node.Right = &TreeNode{Val: ints[i]}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}

func main() {
	fmt.Println(findSecondMinimumValue(Ints2TreeNode([]int{1, 1, 3, 1, 1, 3, 4, 3, 1})))
	fmt.Println(findSecondMinimumValue(Ints2TreeNode([]int{1, 1, 3, 1, 1, 3, 4, 3, 1, 1, 1, 3, 8, 4, 8, 3, 3, 1, 6, 2, 1})))
	fmt.Println(findSecondMinimumValue(Ints2TreeNode([]int{2, 2, 5, null, null, 5, 7})))
	fmt.Println(findSecondMinimumValue(Ints2TreeNode([]int{2, 2, 2})))
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findSecondMinimumValue(root *TreeNode) int {
	if root.Left != nil && root.Right != nil {
		if root.Left.Val > root.Right.Val {
			v := findSecondMinimumValue(root.Right)
			if v > root.Left.Val || v == -1 {
				return root.Left.Val
			} else if v <= root.Left.Val {
				return v
			}
		} else if root.Left.Val < root.Right.Val {
			v := findSecondMinimumValue(root.Left)
			if v > root.Right.Val || v == -1 {
				return root.Right.Val
			} else if v <= root.Right.Val {
				return v
			}
		} else {
			v1 := findSecondMinimumValue(root.Left)
			v2 := findSecondMinimumValue(root.Right)
			if v1 < v2 || v2 == -1 {
				return v1
			}
			if v1 > v2 || v1 == -1 {
				return v2
			}
			return -1
		}
	} else if root.Right != nil {
		return findSecondMinimumValue(root.Right)
	} else if root.Left != nil {
		return findSecondMinimumValue(root.Left)
	}
	return -1
}
