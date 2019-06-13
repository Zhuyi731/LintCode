package main 
import (
	"fmt"
	"strings"
	"strconv"
)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

 func main(){
	a:=swapPairs(CreateListNode("1->2->3->4"))
 }

 func swapPairs(head *ListNode) *ListNode {
	var result = head 
	var cur = *head

	for ;cur != nil && cur.Next != nil ;{
		next := &cur.Next.Next
		cur.Next.Next = &cur
		cur.Next = next  
		cur = cur.Next 
	}	

	return result
}

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
// Given a linked list, swap every two adjacent nodes and return its head.

// You may not modify the values in the list's nodes, only nodes itself may be changed.

 

// Example:

// Given 1->2->3->4, you should return the list as 2->1->4->3.