package main

import "fmt"

func main() {
	fmt.Println(searchMatrix([][]int{{-5}}, -5))
	fmt.Println(searchMatrix([][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}, 5))
	fmt.Println(searchMatrix([][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}, 20))
}

func searchMatrix(matrix [][]int, target int) bool {
	w := len(matrix)
	h := len(matrix[0])

	startX := 0
	startY := h - 1
	for startX < w && startY >= 0 {
		if matrix[startX][startY] > target {
			startY--
		} else if matrix[startX][startY] < target {
			startX++
		} else {
			return true
		}
	}
	return false
}
