package main

import "fmt"

func main() {
	fmt.Println(maxIncreaseKeepingSkyline([][]int{{3, 0, 8, 4}, {2, 4, 5, 7}, {9, 2, 6, 3}, {0, 3, 1, 0}}))
	fmt.Println(maxIncreaseKeepingSkyline([][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}))
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func maxIncreaseKeepingSkyline(grid [][]int) int {
	//
	result := 0
	m := len(grid)
	if len(grid[0]) > m {
		m = len(grid[0])
	}
	maxCache := make([][]int, m)

	for i := 0; i < m; i++ {
		maxCache[i] = []int{-1, -1}
	}
	for i, row := range grid {
		for j, v := range row {
			mc := 0
			mr := 0
			if maxCache[i][0] != -1 { // 横向最高
				mr = maxCache[i][0]
			} else {
				for k := 0; k < len(grid[0]); k++ {
					if grid[i][k] > mr {
						mr = grid[i][k]
					}
				}
				maxCache[i][0] = mr
			}
			if maxCache[j][1] != -1 {
				mc = maxCache[j][1]
			} else {
				for k := 0; k < len(grid); k++ {
					if grid[k][j] > mc {
						mc = grid[k][j]
					}
				}
				maxCache[j][1] = mc
			}
			result += min(mc, mr) - v
		}
	}
	return result
}
