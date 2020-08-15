package main

import "fmt"

func main() {
	rotate([][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16},
	})
	rotate([][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	})

}

func rotate(matrix [][]int) {
	n := len(matrix)
	// 首先对角线翻转
	for i := 0; i < n; i++ {
		for j := 0; j < n-i; j++ {
			matrix[i][j], matrix[n-j-1][n-i-1] = matrix[n-j-1][n-i-1], matrix[i][j]
		}
	}
	// print(matrix)

	// 然后上下翻转

	for i := 0; i < n/2; i++ {
		for j := 0; j < n; j++ {
			// print(matrix)
			matrix[i][j], matrix[n-i-1][j] = matrix[n-i-1][j], matrix[i][j]
		}
	}

	fmt.Println(matrix)
}

func print(matrix [][]int) {
	for _, val := range matrix {
		for _, val2 := range val {
			fmt.Printf("%d ", val2)
		}
		fmt.Println()
	}
}
