/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
package main

import (
	// "/Users/zhuyi/Desktop/code/temp/LintCode/go_utils/node"
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var NULL = -1 << 63

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
	fmt.Println(minDiffInBST(Ints2TreeNode([]int{27, NULL, 34, NULL, 58, 50, NULL, 44})))                                 //6
	fmt.Println(minDiffInBST(Ints2TreeNode([]int{56, 42, 77, 23, 51, 75, 90, NULL, NULL, NULL, NULL, 67, NULL, 78, 99}))) //6
	fmt.Println(minDiffInBST(Ints2TreeNode([]int{90, 69, NULL, 49, 89, NULL, 52})))                                       //6
	fmt.Println(minDiffInBST(Ints2TreeNode([]int{27, NULL, 34, NULL, 58, 50, NULL, 44})))                                 //6
	fmt.Println(minDiffInBST(Ints2TreeNode([]int{1, 0, 48, NULL, NULL, 12, 49})))                                         //1
	fmt.Println(minDiffInBST(Ints2TreeNode([]int{4, 2, 6, 1, 3})))                                                        //1
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

var min = 1000000

func minDiffInBST(root *TreeNode) int {
	return f(root, min)
}

func f(root *TreeNode, parent int) int {
	l := min
	lc := min
	r := min
	rc := min
	rp := min
	lp := min
	if root.Left != nil {
		l = abs(root.Left.Val - root.Val)
		lc = f(root.Left, root.Val)
		lp = abs(parent - root.Left.Val)
	}
	if root.Right != nil {
		r = abs(root.Right.Val - root.Val)
		rc = f(root.Right, root.Val)
		rp = abs(parent - root.Right.Val)
	}

	if l < min {
		min = l
	}

	if lc < min {
		min = lc
	}
	if lp < min {
		min = lp
	}
	if r < min {
		min = r
	}
	if rc < min {
		min = rc
	}
	if rp < min {
		min = rp
	}

	return min
}
