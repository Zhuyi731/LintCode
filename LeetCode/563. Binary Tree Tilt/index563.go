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
		if res[i] != "nil" && res[i] != "null" {
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
			if rightIndex < length && treeNode[rightIndex] != nil {
				treeNode[i].Right = treeNode[rightIndex]
			}
		}
	}
	return treeNode[0]
}

func main() {
	fmt.Println(findTilt(CreateTestData("[]")))
	fmt.Println(findTilt(CreateTestData("[1]")))
	fmt.Println(findTilt(CreateTestData("[1,2]")))
	fmt.Println(findTilt(CreateTestData("[1,null,2]")))
	fmt.Println(findTilt(CreateTestData("[4,2,9,3,5,null,7]")))
	fmt.Println(findTilt(CreateTestData("[1,2,3]")))
	fmt.Println(findTilt(CreateTestData("[21,7,14,1,1,2,2,3,3]")))
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findTilt(root *TreeNode) int {
	if root == nil {
		return 0
	}
	_, val := dfs(root)
	return val
}

func dfs(root *TreeNode) (sum, diffSum int) {
	sum = 0
	diff := 0
	diffSum = 0
	if root.Left != nil {
		sumL, diffSL := dfs(root.Left)
		sum += sumL
		diffSum += diffSL
		diff += sumL
	}
	if root.Right != nil {
		sumR, diffSR := dfs(root.Right)
		sum += sumR
		diffSum += diffSR
		diff -= sumR
	}

	return sum + root.Val, diffSum + abs(diff)
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
