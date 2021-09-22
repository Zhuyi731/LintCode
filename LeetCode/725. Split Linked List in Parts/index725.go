package main

func main() {
	splitListToParts(CreateListNodeBySlice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}), 5)
	splitListToParts(CreateListNodeBySlice([]int{1, 2, 3}), 5)
}

// ListNode 链表结构
type ListNode struct {
	Val  int       //当前节点的值
	Next *ListNode // 下一节点指针
}

func CreateListNodeBySlice(nodes []int) *ListNode {
	nodeList := ListNode{}
	curNode := &nodeList
	for index, value := range nodes {
		curNode.Val = value
		if index != len(nodes)-1 {
			tempNode := &ListNode{}
			curNode.Next = tempNode
			curNode = tempNode
		}
	}

	return &nodeList
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func splitListToParts(head *ListNode, k int) []*ListNode {
	l := 0
	tp := head
	for tp != nil {
		tp = tp.Next
		l++
	}

	result := make([]*ListNode, k)
	lenR := l / k
	restR := l % k
	curIndex := 0
	curNode := head
	tmpNode := head
	ct := 0

	result[curIndex] = curNode
	for i := 1; i < l; i++ {
		if curIndex < restR {
			if ct < lenR {
				tmpNode = tmpNode.Next
				ct++
			} else {
				curNode = tmpNode.Next
				tmpNode.Next = nil
				curIndex++
				result[curIndex] = curNode
				tmpNode = curNode
				ct = 0
			}
		} else {
			if ct < lenR-1 {
				tmpNode = tmpNode.Next
				ct++
			} else {
				if tmpNode == nil {
					curIndex++
					result[curIndex] = nil
				} else {
					curNode = tmpNode.Next
					tmpNode.Next = nil
					curIndex++
					result[curIndex] = curNode
					tmpNode = curNode
				}
				ct = 0
			}
		}
	}
	return result
}
