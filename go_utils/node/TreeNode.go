package node

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

const spce = "      "

// 打印树
func Print(root *TreeNode) {
	treeNodePrint(root, 0)
}

func treeNodePrint(node *TreeNode, deep int) {
	if node == nil {
		printSpace(deep)
		fmt.Println("#")
		return
	}
	treeNodePrint(node.Right, deep+1)
	printSpace(deep)
	printNode(node.Val)
	treeNodePrint(node.Left, deep+1)
}

func printSpace(count int) {
	for i := 0; i < count; i++ {
		fmt.Printf(spce)
	}
}

func printNode(val int) {
	var buffer strings.Builder
	temp := strconv.Itoa(val)
	buffer.WriteString(temp)
	buffer.WriteString("<")
	spceNum := len(spce) - buffer.Len()
	for i := 0; i < spceNum; i++ {
		buffer.WriteString(" ")
	}
	fmt.Println(buffer.String())
}
