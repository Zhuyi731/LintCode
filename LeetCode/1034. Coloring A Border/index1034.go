package main

import "fmt"

func main() {
	fmt.Println(colorBorder([][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}}, 1, 1, 2))
	fmt.Println(colorBorder([][]int{{1, 1},
		{1, 2}}, 0, 0, 3))
	fmt.Println(colorBorder([][]int{{1, 2, 2},
		{2, 3, 2}}, 0, 1, 3))
}

func colorBorder(grid [][]int, row int, col int, color int) [][]int {
	originColor := grid[row][col]

	result := make([][]int, len(grid))
	for i, value := range grid {
		result[i] = make([]int, len(grid[i]))
		for j, v := range value {
			result[i][j] = v
		}
	}
	travelArr := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	var checkOk = func(r, c int) bool {
		return !(r >= len(grid) || r < 0 || c >= len(grid[0]) || c < 0)
	}
	var dfs func(r, c int)
	dfs = func(r, c int) {
		if grid[r][c] == -1 {
			return
		}
		grid[r][c] = -1
		for _, v := range travelArr {
			i, j := v[0], v[1]
			if checkOk(r+i, c+j) && (grid[r+i][c+j] == originColor || grid[r+i][c+j] == -1) { // 说明该点需要遍历
				dfs(r+i, c+j)
			} else {
				result[r][c] = color
			}
		}

	}
	dfs(row, col)
	return result
}
