package main

import "fmt"

func main() {
	fmt.Println(fourSumCount(
		[]int{1, 2},
		[]int{-2, -1},
		[]int{-1, 2},
		[]int{0, 2},
	))
}

func fourSumCount(A []int, B []int, C []int, D []int) int {
	hashMapper1 := make(map[int]int)

	for i := 0; i < len(A); i++ {
		for j := 0; j < len(B); j++ {
			sum := A[i] + B[j]
			hashMapper1[sum] = hashMapper1[sum] + 1
		}
	}

	result := 0
	for i := 0; i < len(C); i++ {
		for j := 0; j < len(D); j++ {
			result += hashMapper1[-C[i]-D[j]]
		}
	}

	return result
}
