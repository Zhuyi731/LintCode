package main

import "fmt"

func main() {

	fmt.Println(searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 16))
	fmt.Println(searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 60))
	fmt.Println(searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 5))
	fmt.Println(searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 1))

	fmt.Println(searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}, {61, 69, 76, 88}}, 16))
	fmt.Println(searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}, {61, 69, 76, 88}}, 1))
	fmt.Println(searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}, {61, 69, 76, 88}}, 34))
	fmt.Println(searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}, {61, 69, 76, 88}}, 5))
	fmt.Println(searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}, {61, 69, 76, 88}}, 88))

	fmt.Println("/*************/")

	fmt.Println(searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, -2))
	fmt.Println(searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 2))
	fmt.Println(searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 13))
	fmt.Println(searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 26))
	fmt.Println(searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 126))
	fmt.Println(searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}, {61, 69, 76, 88}}, -1))
	fmt.Println(searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}, {61, 69, 76, 88}}, 2))
	fmt.Println(searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}, {61, 69, 76, 88}}, 12))
	fmt.Println(searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}, {61, 69, 76, 88}}, 32))
	fmt.Println(searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}, {61, 69, 76, 88}}, 68))
	fmt.Println(searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}, {61, 69, 76, 88}}, 268))
}

func searchMatrix(matrix [][]int, target int) bool {
	//  confirm row
	l, r := 0, len(matrix)-1

	for {
		if l >= r {
			break
		}

		m := (l + r) / 2

		if target > matrix[m][0] {
			l = m + 1
		} else if target < matrix[m][0] {
			r = m - 1
		} else {
			return true
		}
	}

	line := l
	if target < matrix[line][0] {
		line--
	}
	if line < 0 {
		return false
	}

	l, r = 0, len(matrix[0])-1
	for {
		if l >= r {
			break
		}

		m := (l + r) / 2

		if target > matrix[line][m] {
			l = m + 1
		} else if target == matrix[line][m] {
			return true
		} else {
			r = m - 1
		}
	}

	return target == matrix[line][r]
}
