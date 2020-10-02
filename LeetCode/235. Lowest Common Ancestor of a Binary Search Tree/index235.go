package main

import (
	"fmt"
	"strconv"
	"strings"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

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
	// fmt.Println(lowestCommonAncestor(CreateTestData("[6,2,8,0,4,7,9,nil,nil,3,5]"), CreateTestData("[2]"), CreateTestData("[8]")).Val)
	fmt.Println(lowestCommonAncestor(CreateTestData("[6,2,8,0,4,7,9,nil,nil,3,5]"), CreateTestData("[2]"), CreateTestData("[4]")).Val)
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var v1, v2 int
	if p.Val > q.Val {
		v1 = q.Val
		v2 = p.Val
	} else {
		v1 = p.Val
		v2 = q.Val
	}
	for {
		if root.Val == v2 || root.Val == v1 {
			return root
		} else if root.Val > v1 && root.Val < v2 {
			return root
		} else if root.Val > v1 && root.Val > v2 {
			root = root.Left
		} else if root.Val < v1 && root.Val < v2 {
			root = root.Right
		}
	}
	return &TreeNode{Val: 0}
}
