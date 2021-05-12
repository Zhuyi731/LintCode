package main

import "fmt"

func main() {
	fmt.Println(xorQueries([]int{1, 3, 4, 8}, [][]int{{0, 1}, {1, 2}, {0, 3}, {3, 3}}))
	fmt.Println(xorQueries([]int{4, 8, 2, 10}, [][]int{{2, 3}, {1, 3}, {0, 0}, {0, 3}}))
}

func xorQueries(arr []int, queries [][]int) []int {
	sum := make([]int, len(arr)+1)
	if len(arr) == 0 {
		return []int{}
	}
	sum[0] = 0
	for i, v := range arr {
		sum[i+1] = sum[i] ^ v
	}

	result := make([]int, len(queries))
	for i, v := range queries {
		result[i] = sum[v[0]] ^ sum[v[1]+1]
	}
	return result
}
