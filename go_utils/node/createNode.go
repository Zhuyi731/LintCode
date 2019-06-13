package node

import (
	"strings"
	"strconv"
)

type ListNode struct {
	Val int
	Next *ListNode
}

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