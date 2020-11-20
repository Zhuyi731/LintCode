package main

import (
	"fmt"
	"strconv"
	"strings"
)

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

type ListNode struct {
	Val  int       //当前节点的值
	Next *ListNode // 下一节点指针
}

func main() {
	insertionSortList(CreateListNode("4->2->1->3"))
	insertionSortList(CreateListNode("-1->5->3->4->0"))

	insertionSortList(CreateListNode("0"))
	insertionSortList(CreateListNode("1->1"))
	insertionSortList(CreateListNode("1->1->1"))
	insertionSortList(CreateListNode("1->2->3"))
	insertionSortList(CreateListNode("3->2->1"))
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

func insert(newNodeList *ListNode, node *ListNode) {
	currentNewNode := newNodeList
	for {
		if currentNewNode.Next != nil && currentNewNode.Next.Val > node.Val {
			//插入当前位置
			tempNode := currentNewNode.Next
			currentNewNode.Next = node
			node.Next = tempNode
			break
		} else if currentNewNode.Next == nil {
			currentNewNode.Next = node
			node.Next = nil
			break
		} else {
			currentNewNode = currentNewNode.Next
		}
	}
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }

 */
func insertionSortList(head *ListNode) *ListNode {
	newNodeList := ListNode{
		Next: nil,
		Val:  -0x7ffffff,
	}
	currentNode := head

	for {
		if currentNode == nil {
			break
		}
		next := currentNode.Next
		insert(&newNodeList, currentNode)
		currentNode = next
	}
	PrintListNode(newNodeList.Next)
	return newNodeList.Next
}
