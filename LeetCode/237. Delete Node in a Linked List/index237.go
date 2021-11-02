/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteNode(node *ListNode) {
	t := node
	t.Val = t.Next.Val
	t.Next = t.Next.Next
	node = t
}


