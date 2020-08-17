package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func createTreeNodeByArr(arr []int) *TreeNode {
	if len(arr) == 0 {
		return nil
	}
	currentNodeList := make([]*TreeNode, len(arr))
	for index, val := range arr {
		var tempNode TreeNode
		if index != 0 {
			parent := (index - 1) / 2
			isLeft := (index-1)%2 == 0
			if val == -1 {
				if isLeft {
					currentNodeList[parent].Left = nil
				} else {
					currentNodeList[parent].Right = nil
				}
			} else {
				tempNode = TreeNode{
					Val: val,
				}
				if isLeft {
					currentNodeList[parent].Left = &tempNode
				} else {
					currentNodeList[parent].Right = &tempNode
				}
			}

		}

		currentNodeList[index] = &tempNode
	}
	return currentNodeList[0]
}

func main() {
	fmt.Println(isBalanced(createTreeNodeByArr([]int{1, 2, 2, 3, -1, -1, 3, 4, -1, -1, 4})))
	fmt.Println(isBalanced(createTreeNodeByArr([]int{3, 9, 20, -1, -1, 15, 7})))
	fmt.Println(isBalanced(createTreeNodeByArr([]int{1, 2, 2, 3, 3, -1, -1, 4, 4})))
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	left := getLen(root.Left, 0)
	right := getLen(root.Right, 0)

	diff := left - right
	if diff > 1 || diff < -1 || !isBalanced(root.Left) || !isBalanced(root.Right) {
		return false
	}
	return true

}

func getLen(node *TreeNode, currentDepth int) int {
	var left, right int
	left = 0
	right = 0
	if node == nil {
		return currentDepth
	}
	if node.Left == nil && node.Right == nil {
		return currentDepth + 1
	}
	if node.Left != nil {
		left = getLen(node.Left, currentDepth+1)
	}
	if node.Right != nil {
		right = getLen(node.Right, currentDepth+1)
	}
	if left > right {
		return left
	}
	return right
}
