package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

const NULL = 0X7FFFFFFF
const null = 0x7FFFFFFF

func main() {
	fmt.Println(isCousins(Ints2TreeNode([]int{1, 2, 3, 4}), 3, 4))
	fmt.Println(isCousins(Ints2TreeNode([]int{1, 2, 3, null, 4, null, 5}), 5, 4))
	fmt.Println(isCousins(Ints2TreeNode([]int{1, 2, 3, null, 4}), 2, 3))
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

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

func isCousins(root *TreeNode, x int, y int) bool {
	var p1, p2, d1, d2 int

	var dfs func(r *TreeNode, depth int, parent int)
	dfs = func(r *TreeNode, depth int, parent int) {
		if x == r.Val {
			p1 = parent
			d1 = depth
		}
		if y == r.Val {
			p2 = parent
			d2 = depth
		}
		if r.Left != nil {
			dfs(r.Left, depth+1, r.Val)
		}
		if r.Right != nil {
			dfs(r.Right, depth+1, r.Val)
		}
	}

	dfs(root, 0, -1)
	return d1 == d2 && p1 != p2
}
