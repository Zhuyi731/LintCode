package node

import (
	"fmt"
	"strconv"
	"strings"
)

// ListNode 链表结构
type ListNode struct {
	Val  int       //当前节点的值
	Next *ListNode // 下一节点指针
}

func PrintListNode(node *ListNode) {
	bak := node
	isFirst := true
	for {
		if bak == nil {
			fmt.Println("")
			return
		}
		if isFirst {
			fmt.Print(bak.Val)
			isFirst = false
		} else {
			fmt.Print("->", bak.Val)
		}
		bak = bak.Next
	}
}

// CreateListNode 通过string 创建链表结构
func CreateListNode(nodeString string) *ListNode {
	nodeList := ListNode{}
	node := strings.Split(nodeString, "->")
	curNode := &nodeList
	for index, value := range node {
		curNode.Val, _ = strconv.Atoi(value)
		if index != len(node)-1 {
			tempNode := &ListNode{}
			curNode.Next = tempNode
			curNode = tempNode
		}
	}

	return &nodeList
}

func CreateListNodeBySlice(nodes []int) *ListNode {
	nodeList := ListNode{}
	curNode := &nodeList
	for index, value := range nodes {
		curNode.Val = value
		if index != len(node)-1 {
			tempNode := &ListNode{}
			curNode.Next = tempNode
			curNode = tempNode
		}
	}

	return &nodeList
}

// Node 无向图结构
type Node struct {
	Val       int     //当前节点值
	Neighbors []*Node //相邻节点
}

// CreateNodeBySlice 通过数组创建无向图
func CreateNodeBySlice(arr [][]int) *Node {
	nodeList := make([]Node, len(arr))

	for i := 0; i < len(arr); i++ {
		nodeList[i] = Node{
			Val: i + 1,
		}
		tempNodeList := make([]*Node, len(arr[i]))
		for j := 0; j < len(arr[i]); j++ {
			currentIndex := arr[i][j] - 1
			if nodeList[currentIndex].Val != currentIndex+1 { // 说明没有创建过
				nodeList[currentIndex].Val = currentIndex + 1
			}

			tempNodeList[j] = &nodeList[currentIndex]
		}
		nodeList[i].Neighbors = tempNodeList
	}
	return &nodeList[0]
}
