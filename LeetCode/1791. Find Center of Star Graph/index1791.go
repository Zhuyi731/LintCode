package main

import "fmt"

func main() {
	fmt.Println(findCenter([][]int{{1, 2}, {2, 3}, {4, 2}}))
}

func findCenter(edges [][]int) int {
	if edges[0][1] == edges[1][0] || edges[0][1] == edges[1][1] {
		return edges[0][1]
	}
	return edges[0][0]
}
