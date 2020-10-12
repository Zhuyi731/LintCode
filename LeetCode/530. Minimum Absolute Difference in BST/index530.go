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
			if rightIndex < length && treeNode[rightIndex] != nil {
				treeNode[i].Right = treeNode[rightIndex]
			}
		}
	}
	return treeNode[0]
}

func main() {
	tree := CreateTestData("[334,277,507,nil,285,nil,678]")
	fmt.Println(getMinimumDifference(tree))

	tree = CreateTestData("[1,nil,3,nil,nil,2]")
	fmt.Println(getMinimumDifference(tree))
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getMinimumDifference(root *TreeNode) int {
	arr := make([]int, 0)
	return solution(root, &arr)
}

const INT_MAX = int(^uint(0) >> 1)

func solution(root *TreeNode, currentArr *[]int) (currentMin int) {
	currentVal := root.Val
	for index, val := range *currentArr {
		if currentVal < val {
			*currentArr = append((*currentArr)[:index], append([]int{root.Val}, (*currentArr)[index:]...)...)
			currentMin = (*currentArr)[index+1] - (*currentArr)[index]
			if index > 0 && (*currentArr)[index]-(*currentArr)[index-1] < currentMin {
				currentMin = (*currentArr)[index] - (*currentArr)[index-1]
			}
			leftMin := INT_MAX
			rightMin := INT_MAX
			if root.Left != nil {
				leftMin = solution(root.Left, currentArr)
				if leftMin < currentMin {
					currentMin = leftMin
				}
			}
			if root.Right != nil {
				rightMin = solution(root.Right, currentArr)
				if rightMin < currentMin {
					currentMin = rightMin
				}
			}
			return currentMin
		}
	}
	*currentArr = append(*currentArr, root.Val)
	if len(*currentArr) > 1 {
		currentMin = (*currentArr)[len(*currentArr)-1] - (*currentArr)[len(*currentArr)-2]
	} else {
		currentMin = INT_MAX
	}
	leftMin := INT_MAX
	rightMin := INT_MAX
	if root.Left != nil {
		leftMin = solution(root.Left, currentArr)
		if leftMin < currentMin {
			currentMin = leftMin
		}
	}
	if root.Right != nil {
		rightMin = solution(root.Right, currentArr)
		if rightMin < currentMin {
			currentMin = rightMin
		}
	}
	return currentMin

}
