/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	nodes := []*Node{root}
	nodesN := make([]*Node, 0)
	for {
		for i,v := nodes {
			if v.Left != nil {
				nodesN = append(nodesN,v.Left)
			}
			if v.Right != nil {
				nodesN = append(nodesN,v.Right)
			}
		}
		ln := len(nodesN)
		if ln ==0 {
			break
		}
		for i:=1;i<ln;i++{
			nodesN[i-1].Next = nodesN[i]
		}
		nodes = nodesN 
		nodesN = make([]*Node,0)
	}
}