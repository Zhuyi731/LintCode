package main

import (
	"strconv"
	"strings"
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// printNode(mergeTwoLists(createListNode("1->2->4"), createListNode("1->3->4")))
	printNode(mergeTwoLists(createListNode("1->1->1"), createListNode("1->3->4")))
	printNode(mergeTwoLists(createListNode("2->3->4"), createListNode("1->3->4")))
	printNode(mergeTwoLists(createListNode("1->6->10"), createListNode("1->3->4")))
	printNode(mergeTwoLists(createListNode("1->2->3"), createListNode("1->1->1")))
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	//让l1的第一个节点总是最小
	if (*l1).Val > (*l2).Val {
		temp := &l1
		l1 = l2
		l2 = *temp
	}
	 ret := l1 

	curL1Node := l1
	curL2Node := l2

	var tl2,tl1 *ListNode

	for {
		if curL2Node == nil || curL2Node.Next == nil {
			break
		}

		for {
			if curL1Node.Next == nil{
				curL1Node.Next = curL2Node
				for {
					if curL2Node != nil{
						curL2Node=curL2Node.Next
					}else{
						break
					}
				}
				break
			}

			if curL1Node.Val <= curL2Node.Val &&  curL1Node.Next.Val >curL2Node.Val {
				tl1 = curL1Node.Next
				curL1Node.Next = curL2Node
				tl2 = curL2Node.Next 
				curL1Node = curL1Node.Next 
				curL2Node = tl2
				curL1Node.Next = tl1
				curL1Node = curL1Node.Next
				printNode(ret)
				printNode(curL1Node)
				printNode(curL2Node)
			}else{
				curL1Node = curL1Node.Next
			}
		}
		if curL2Node != nil{
			curL2Node = curL2Node.Next
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

func printNode (node *ListNode){
	if node != nil {
		fmt.Print(node.Val)
		node = node.Next
	}else{
		return 
	}
	for{
		if node != nil {
			fmt.Print("->",node.Val)
			node = node.Next
		} else{
			break
		}
	}
	fmt.Print("\n")
}
// Merge two sorted linked lists and return it as a new list. The new list should be made by splicing together the nodes of the first two lists.

// Example:

// Input: 1->2->4, 1->3->4
// Output: 1->1->2->3->4->4
