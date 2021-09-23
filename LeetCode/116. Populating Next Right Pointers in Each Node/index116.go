/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func patchNode(l,r *Node) {
	for ;l!=nil; {
		l = l.Right 
		r = r.Left 
		if l.Next ==nil {
			l.Next = r
		}
	}
}

func connect(root *Node) *Node {
	if root == nil return 
	if root.Left != nil  {
		root.Left.Next = root.Right
		connect(root.Left)
		connect(root.Right)
		patchNode(root.Left,root.Right)
		return root.Right
	} else  {
		return root
	}
}