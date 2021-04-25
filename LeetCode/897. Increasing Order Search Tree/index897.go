/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
package main

import "fmt"

const NULL = 0xfffff
const null = NULL

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

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
	fmt.Println(increasingBST(Ints2TreeNode([]int{5, 1, 7})))
	fmt.Println(increasingBST(Ints2TreeNode([]int{5, 3, 6, 2, 4, null, 8, 1, null, null, null, 7, 9})))
}

func increasingBST(root *TreeNode) *TreeNode {
	res := TreeNode{}

	var toIncreasingBst func(root *TreeNode, result *TreeNode) *TreeNode
	toIncreasingBst = func(root *TreeNode, result *TreeNode) *TreeNode {
		if root.Left != nil {
			result = toIncreasingBst(root.Left, result)
		}
		result.Right = &TreeNode{
			Val:   root.Val,
			Left:  nil,
			Right: nil,
		}
		result = result.Right
		if root.Right != nil {
			result = toIncreasingBst(root.Right, result)
		}
		return result
	}

	toIncreasingBst(root, &res)

	return res.Right
}
