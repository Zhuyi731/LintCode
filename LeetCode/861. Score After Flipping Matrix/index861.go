package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(matrixScore([][]int{{0, 0, 1, 1}, {1, 0, 1, 0}, {1, 1, 0, 0}}))
}

func matrixScore(A [][]int) int {
	if len(A) == 0 {
		return 0
	}
	for _, item := range A {
		if item[0] == 0 {
			for index, val := range item {
				item[index] = (val + 1) % 2
			}
		}
	}

	lens := len(A)
	width := len(A[0])

	result := math.Pow(2, float64(width-1)) * float64(lens)

	for i := 1; i < width; i++ {
		// 每一列计数
		ct := 0
		for j := 0; j < lens; j++ {
			if A[j][i] == 1 {
				ct++
			}
		}
		if ct*2 < lens {
			ct = lens - ct
		}
		result += math.Pow(2, float64(width-i-1)) * float64(ct)
	}

	return int(result)
}
