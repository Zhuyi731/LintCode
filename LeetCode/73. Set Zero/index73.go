package main

import "fmt"

func main() {
	setZeroes([][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}})
	setZeroes([][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}})
}

func setZeroes(matrix [][]int) {
	rows := []int{}
	cols := []int{}

	r := len(matrix)
	c := len(matrix[0])

	for i, arr := range matrix {
		for j, v := range arr {
			if v == 0 {
				rows = append(rows, i)
				cols = append(cols, j)
			}
		}
	}

	for _, v := range rows {
		for i := 0; i < c; i++ {
			matrix[v][i] = 0
		}
	}

	for _, v := range cols {
		for i := 0; i < r; i++ {
			matrix[i][v] = 0
		}
	}

	fmt.Print(matrix)
}
