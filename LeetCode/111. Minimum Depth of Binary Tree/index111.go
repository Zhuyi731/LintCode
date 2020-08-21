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
	currentNodeList := make([]TreeNode, len(arr))
	for index, val := range arr {
		tempNode := TreeNode{
			Val: val,
		}
		if index != 0 {
			parent := (index - 1) / 2
			isLeft := (index-1)%2 == 0
			if isLeft {
				currentNodeList[parent].Left = &tempNode
			} else {
				currentNodeList[parent].Right = &tempNode
			}
		}
	}
	return &currentNodeList[0]
}

func main() {
	/**
	*
	**/
	node := createTreeNodeByArr([]int{3, 9, 20, -1, -1, 15, 7})
	fmt.Println(minDepth(node))
	node = createTreeNodeByArr([]int{1, 2, 5, 3, -1, 6, -1, 4, -1, 7})
	fmt.Println(minDepth(node))
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
	return getDepth(root, 1)
}

func getDepth(root *TreeNode, currentDepth int) int {
	if root == nil {
		return currentDepth - 1
	}
	// fmt.Println(root.Left, root.Right, currentDepth)
	if root.Left == nil && root.Right == nil {
		return currentDepth
	}
	if root.Left == nil {
		return getDepth(root.Right, currentDepth+1)
	} else if root.Right == nil {
		return getDepth(root.Left, currentDepth+1)
	} else {
		left := getDepth(root.Left, currentDepth+1)
		right := getDepth(root.Right, currentDepth+1)
		if left > right {
			return right
		}
		return left
	}
}
