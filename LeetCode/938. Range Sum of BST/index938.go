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

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

const NULL = 0x7fffff
const null = 0x7fffff

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
	fmt.Println(rangeSumBST(Ints2TreeNode([]int{10, 5, 15, 3, 7, null, 18}), 7, 15))
	fmt.Println(rangeSumBST(Ints2TreeNode([]int{10, 5, 15, 3, 7, 13, 18, 1, null, 6}), 6, 10))
}

func rangeSumBST(root *TreeNode, low int, high int) int {
	sum := 0
	var travel func(*TreeNode)
	travel = func(root *TreeNode) {
		// do something
		if root.Val <= high && root.Val >= low {
			sum += root.Val
			if root.Left != nil {
				travel(root.Left)
			}
			if root.Right != nil {
				travel(root.Right)
			}
		} else if root.Val < low {
			if root.Right != nil {
				travel(root.Right)
			}
		} else {
			if root.Left != nil {
				travel(root.Left)
			}
		}
	}
	travel(root)
	return sum
}
