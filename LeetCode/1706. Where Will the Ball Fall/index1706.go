package main

import "fmt"

func main() {
	fmt.Println(findBall([][]int{{1, 1, 1, 1, 1, 1}, {-1, -1, -1, -1, -1, -1}, {1, 1, 1, 1, 1, 1}, {-1, -1, -1, -1, -1, -1}}))
	fmt.Println(findBall([][]int{{1, 1, 1, -1, -1}, {1, 1, 1, -1, -1}, {-1, -1, -1, 1, 1}, {1, 1, 1, 1, -1}, {-1, -1, -1, -1, -1}}))
}

func findBall(grid [][]int) []int {
	ans := make([]int, len(grid[0]))

	check := func(currentPos, y int) bool {
		return currentPos+grid[y][currentPos] >= 0 && currentPos+grid[y][currentPos] < len(grid[0]) && grid[y][currentPos] == grid[y][currentPos+grid[y][currentPos]]
	}

	for i := 0; i < len(grid[0]); i++ {
		// 检测第i个是否通过
		currentPos := i
		for y := 0; y < len(grid); y++ {
			if check(currentPos, y) {
				currentPos += grid[y][currentPos]
			} else {
				ans[i] = -1
				break
			}
		}
		if ans[i] != -1 {
			ans[i] = currentPos
		}
	}
	return ans
}
