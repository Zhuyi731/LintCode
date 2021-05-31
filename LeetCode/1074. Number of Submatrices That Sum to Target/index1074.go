package main

import "fmt"

func main() {
	fmt.Println(numSubmatrixSumTarget([][]int{{0, 0, 0, 1, 1}, {1, 1, 1, 1, 1}, {0, 1, 0, 0, 0}, {0, 1, 0, 0, 0}, {1, 1, 1, 1, 0}, {1, 1, 1, 0, 1}}, 0))
	fmt.Println(numSubmatrixSumTarget([][]int{{0, 1, 0}, {1, 1, 1}, {0, 1, 0}}, 0))
	fmt.Println(numSubmatrixSumTarget([][]int{{1, -1}, {-1, 1}}, 0))
	fmt.Println(numSubmatrixSumTarget([][]int{{904}}, 0))
}

func subarraySum(nums []int, k int) int {
	cur, result := 0, 0
	mapper := map[int]int{
		0: 1,
	}
	for _, v := range nums {
		cur = cur + v
		result += mapper[cur-k]
		mapper[cur]++
	}
	return result
}

func numSubmatrixSumTarget(matrix [][]int, target int) int {
	sum := make([][]int, len(matrix)+1)
	sum[0] = make([]int, len(matrix[0])+1)
	for i, arr := range matrix {
		sum[i+1] = make([]int, len(matrix[0])+1)
		for j, v := range arr {
			sum[i+1][j+1] = sum[i][j+1] + sum[i+1][j] - sum[i][j] + v
		}
	}
	// 然后遍历上下边界
	ct := 0

	for iy := 0; iy < len(matrix[0]); iy++ {
		for jy := iy; jy < len(matrix[0]); jy++ {
			//  从 ix 到 iy列 每一列的和
			row := make([]int, len(matrix))
			for i := 0; i < len(matrix); i++ {
				row[i] = sum[i+1][jy+1] - sum[i][jy+1] - sum[i+1][iy] + sum[i][iy]
			}
			ct += subarraySum(row, target)
		}
	}
	return ct
}
