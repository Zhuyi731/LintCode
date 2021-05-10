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
	"fmt"
)

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
	fmt.Println(leafSimilar(Ints2TreeNode([]int{3, 5, 1, 6, 7, 4, 2, null, null, null, null, null, null, 9, 11, null, null, 8, 10}), Ints2TreeNode([]int{3, 5, 1, 6, 2, 9, 8, null, null, 7, 4})))
	fmt.Println(leafSimilar(Ints2TreeNode([]int{3, 5, 1, 6, 2, 9, 8, null, null, 7, 4}), Ints2TreeNode([]int{3, 5, 1, 6, 7, 4, 2, null, null, null, null, null, null, 9, 8})))
	fmt.Println(leafSimilar(Ints2TreeNode([]int{1}), Ints2TreeNode([]int{1})))
	fmt.Println(leafSimilar(Ints2TreeNode([]int{1}), Ints2TreeNode([]int{2})))
	fmt.Println(leafSimilar(Ints2TreeNode([]int{1, 2}), Ints2TreeNode([]int{2, 2})))
	fmt.Println(leafSimilar(Ints2TreeNode([]int{1, 2, 3}), Ints2TreeNode([]int{1, 3, 2})))

}
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	result1 := []int{}
	result2 := []int{}

	var travel func(root *TreeNode, r *[]int)
	travel = func(root *TreeNode, r *[]int) {
		if root.Left == nil && root.Right == nil {
			(*r) = append(*r, root.Val)
			return
		}
		if root.Left != nil {
			travel(root.Left, r)
		}
		if root.Right != nil {
			travel(root.Right, r)
		}
	}
	travel(root1, &result1)
	travel(root2, &result2)
	if len(result1) != len(result2) {
		return false
	}
	for i, v := range result1 {
		if v != result2[i] {
			return false
		}
	}
	return true
}
