package main

import "fmt"

func main() {
	fmt.Println(kthSmallest(Ints2TreeNode([]int{5, 3, 6, 2, 4, null, null, 1}), 3))
	fmt.Println(kthSmallest(Ints2TreeNode([]int{3, 1, 4, null, 2}), 1))
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

const NULL = 0x7FFFFFFF
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

func kthSmallest(root *TreeNode, k int) int {
	val, _ := midTravel(root, k)
	return val
}

func midTravel(root *TreeNode, k int) (int, int) {
	if root.Left != nil {
		val, remainK := midTravel(root.Left, k)
		k = remainK
		if val != -1 {
			return val, k
		}
	}
	k--
	if k == 0 {
		return root.Val, k
	}
	if root.Right != nil {
		val, remainK := midTravel(root.Right, k)
		k = remainK
		if val != -1 {
			return val, k
		}
	}
	return -1, k
}
