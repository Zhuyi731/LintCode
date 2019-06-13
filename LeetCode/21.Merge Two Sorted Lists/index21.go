package main

import (
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	mergeTwoLists(createListNode("1->2->4"), createListNode("1->3->4"))
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	//让l1的第一个节点总是最小
	if (*l1).Val > (*l2).Val {
		temp := &l1
		l1 = l2
		l2 = *temp
	}
	 ret := l1 

	curL1Node := &l1
	curL2Node := &l2

	for {
		if (*curL1Node).Next == nil {
			break
		}
		if (*curL1Node).Next.Val >= (*curL2Node).Val {
			tempNext := (*curL1Node).Next
			(*curL1Node).Next = *curL2Node
			curL2Node = &(*curL2Node).Next
			curL1Node = &(*curL1Node).Next
			(*curL1Node).Next = tempNext
		} else {
			curL1Node = &(*curL1Node).Next
		}
	}

	return ret
}

func createListNode(nodeString string) *ListNode {
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

// Merge two sorted linked lists and return it as a new list. The new list should be made by splicing together the nodes of the first two lists.

// Example:

// Input: 1->2->4, 1->3->4
// Output: 1->1->2->3->4->4
