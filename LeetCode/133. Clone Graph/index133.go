package main

import "fmt"

// Node 无向图结构
type Node struct {
	Val       int     //当前节点值
	Neighbors []*Node //相邻节点
}

// CreateNodeBySlice 通过数组创建无向图
func CreateNodeBySlice(arr [][]int) *Node {
	nodeList := make([]Node, len(arr))

	for i := 0; i < len(arr); i++ {
		nodeList[i] = Node{
			Val: i + 1,
		}
		tempNodeList := make([]*Node, len(arr[i]))
		for j := 0; j < len(arr[i]); j++ {
			currentIndex := arr[i][j] - 1
			if nodeList[currentIndex].Val != currentIndex+1 { // 说明没有创建过
				nodeList[currentIndex].Val = currentIndex + 1
			}

			tempNodeList[j] = &nodeList[currentIndex]
		}
		nodeList[i].Neighbors = tempNodeList
	}
	return &nodeList[0]
}

func main() {
	var test *Node
	test = CreateNodeBySlice([][]int{{2, 4}, {1, 3}, {2, 4}, {1, 3}})
	result := cloneGraph(test)
	fmt.Println(*result)
}

var clonedVal []*Node

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	clonedVal = []*Node{}
	return cloneGraph2(node)
}

func cloneGraph2(node *Node) *Node {
	if result := checkHasCloned(node.Val); result != nil {
		return result
	}
	nodeCopy := Node{
		Val: node.Val,
	}

	// currentIndex := len(clonedVal)
	clonedVal = append(clonedVal, &nodeCopy)

	neighborsCopy := make([]*Node, len(node.Neighbors))
	for index, neibor := range node.Neighbors {
		fmt.Printf("index: %d\n", index)
		neighborsCopy[index] = cloneGraph2(neibor)
	}

	nodeCopy.Neighbors = neighborsCopy

	return &nodeCopy
}

func checkHasCloned(val int) *Node {
	for _, node := range clonedVal {
		if node.Val == val {
			return node
		}
	}
	return nil
}
