package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	node := CreateTestData("[3,9,20,15,7]")
	fmt.Println(averageOfLevels(node))
}

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

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return []float64{0}
	}

	currentNodes := []TreeNode{*root}
	avg := []float64{float64((*root).Val)}
	hasChildren := false
	if root.Left != nil || root.Right != nil {
		hasChildren = true
	}
	for hasChildren {
		tempNodes := make([]TreeNode, 0)
		hasChildren = false
		for _, val := range currentNodes {
			if val.Left != nil {
				hasChildren = true
				tempNodes = append(tempNodes, *val.Left)
			}
			if val.Right != nil {
				hasChildren = true
				tempNodes = append(tempNodes, *val.Right)
			}
		}
		if !hasChildren {
			break
		}
		sum := 0
		for _, val := range tempNodes {
			sum += val.Val
		}
		avg = append(avg, float64(sum)/float64(len(tempNodes)))
		currentNodes = tempNodes
	}
	return avg
}
