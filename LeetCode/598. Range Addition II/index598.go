package main

import "fmt"

func main() {
	fmt.Println(maxCount(3, 3, [][]int{{2, 2}, {3, 3}}))                                                         //4
	fmt.Println(maxCount(3, 3, [][]int{{2, 2}, {3, 3}, {3, 3}, {3, 3}, {2, 2}, {3, 3}, {3, 3}, {3, 3}, {2, 2}})) //4
	fmt.Println(maxCount(3, 3, [][]int{}))                                                                       //9
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxCount(m int, n int, ops [][]int) int {
	minArea := []int{m, n}

	for _, v := range ops {
		x, y := v[0], v[1]
		minArea[0] = min(minArea[0], x)
		minArea[1] = min(minArea[1], y)
	}
	return minArea[0] * minArea[1]
}
