package main

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 生成树结构
func CreateTestData(data string) *TreeNode {
	if data == "[]" {
		return nil
	}
	data = string([]rune(data)[1 : len(data)-1])
	res := strings.Split(data, ",")
	length := len(res)
	treeNode := make([]*TreeNode, length)
	for i := 0; i < length; i++ {
		if res[i] != "nil" {
			val, err := strconv.Atoi(res[i])
			if err != nil {
				panic(err)
			}
			treeNode[i] = &TreeNode{val, nil, nil}
		}
	}
	for i := 0; i < length; i++ {
		if treeNode[i] != nil {
			leftIndex := i*2 + 1
			if leftIndex < length && treeNode[leftIndex] != nil {
				treeNode[i].Left = treeNode[leftIndex]
			}
			rightIndex := leftIndex + 1
			if rightIndex < length && treeNode[leftIndex] != nil {
				treeNode[i].Right = treeNode[rightIndex]
			}
		}
	}
	return treeNode[0]
}

func main() {
	fmt.Println(sumOfLeftLeaves(CreateTestData("[3,9,20,nil,nil,15,7]")))
	fmt.Println(sumOfLeftLeaves(CreateTestData("[3]")))
	fmt.Println(sumOfLeftLeaves(CreateTestData("[]")))
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func sumOfLeftLeaves(root *TreeNode) int {
	return sum(root, true)
}

func sum(root *TreeNode, isLeft bool) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		if isLeft == true {
			return root.Val
		}
		return 0
	}

	return sum(root.Left, true) + sum(root.Right, false)
}
