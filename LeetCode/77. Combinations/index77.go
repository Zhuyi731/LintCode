package main

import "fmt"

func main() {
	fmt.Println(combine(5, 4))
	fmt.Println(combine(6, 6))
	fmt.Println(combine(6, 3))
	fmt.Println(combine(4, 3))
	fmt.Println(combine(4, 2))
	fmt.Println(combine(1, 1))
}

func combine(n int, k int) [][]int {

	result := [][]int{}

	for i := 1; i <= n-k+1; i++ {
		result = append(result, []int{i})
	}

	if k == 1 {
		return result
	}

	curLen := 1
	for {
		newResult := [][]int{}
		for _, v := range result {
			last := v[len(v)-1]
			for i := last + 1; i <= n-k+curLen+1; i++ {
				nv := make([]int, len(v))
				copy(nv, v)
				newV := append(nv, i)
				newResult = append(newResult, newV)
			}
		}
		curLen++
		if curLen == k {
			return newResult
		}
		result = newResult
	}
}
