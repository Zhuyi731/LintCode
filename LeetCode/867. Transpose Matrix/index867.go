package main

func main() {
	transpose([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})
	transpose([][]int{{1, 2, 3}, {4, 5, 6}})
}

func transpose(matrix [][]int) [][]int {
	l := len(matrix)
	if l == 0 {
		return matrix
	}
	h := len(matrix[0])
	tmp := make([][]int, h)

	for i := 0; i < h; i++ {
		tmp[i] = make([]int, l)
		for j := 0; j < l; j++ {
			tmp[i][j] = matrix[j][i]
		}
	}

	return tmp
}
