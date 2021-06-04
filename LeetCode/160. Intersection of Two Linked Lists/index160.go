package main

func main() {
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	ptA, ptB := headA, headB

	for {
		if ptA == ptB {
			break
		}
		if ptA == nil {
			ptA = headB
		} else {
			ptA = ptA.Next
		}
		if ptB == nil {
			ptB = headA
		} else {
			ptB = ptB.Next
		}
	}
	return ptA
}
