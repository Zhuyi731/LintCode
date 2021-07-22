//Definition for a Node.
package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func genNode(a [][]int) *Node {
	nList := make([]Node, len(a))
	for i, v := range a {
		nList[i] = Node{Val: v[0]}
	}
	for i, v := range a {
		if v[1] != -1 {
			nList[i].Random = &nList[v[1]]
		}
		if i != len(a)-1 {
			nList[i].Next = &nList[i+1]
		}
	}
	return &nList[0]
}

func main() {
	//
	copyRandomList(genNode([][]int{{7, -1}, {13, 0}, {11, 4}, {10, 2}, {1, 0}}))
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	nodeList := []Node{}
	ct := 0
	cur := head
	hasMapper := map[*Node]int{}
	for cur != nil {
		newNode := Node{Val: cur.Val}
		nodeList = append(nodeList, newNode)
		hasMapper[cur] = ct
		ct++
		cur = cur.Next
	}
	cur = head
	ct = 0
	for cur != nil {
		if cur.Next != nil {
			nodeList[ct].Next = &nodeList[ct+1]
		}
		if cur.Random != nil {
			nodeList[ct].Random = &nodeList[hasMapper[cur.Random]]
		}
		ct++
		cur = cur.Next
	}

	return &nodeList[0]
}
