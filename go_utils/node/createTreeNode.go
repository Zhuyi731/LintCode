package node

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func createTreeNodeByArr(arr []int) *TreeNode {
	if len(arr) == 0 {
		return nil
	}
	currentNodeList := make([]TreeNode, len(arr))
	for index, val := range arr {
		tempNode := TreeNode{
			Val: val,
		}
		if index != 0 {
			parent := (index - 1) / 2
			isLeft := (index-1)%2 == 0
			if isLeft {
				currentNodeList[parent].Left = &tempNode
			} else {
				currentNodeList[parent].Right = &tempNode
			}
		}
	}
	return &currentNodeList[0]
}
