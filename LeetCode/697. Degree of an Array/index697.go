package main

import "fmt"

func main() {
	fmt.Println(findShortestSubArray([]int{1, 2, 2, 3, 1}))
	fmt.Println(findShortestSubArray([]int{1, 2, 2, 3, 1, 4, 2}))
}

type Ct struct {
	L  int
	R  int
	Ct int
}

func findShortestSubArray(nums []int) int {
	hashMapper := make(map[int]Ct)

	for i, v := range nums {
		if hashMapper[v].Ct == 0 {
			hashMapper[v] = Ct{
				Ct: 1,
				L:  i,
				R:  i,
			}
		} else {
			hashMapper[v] = Ct{
				Ct: hashMapper[v].Ct + 1,
				L:  hashMapper[v].L,
				R:  i,
			}
		}
	}

	max := 0
	minResult := 0
	for _, v := range hashMapper {
		if v.Ct > max {
			max = v.Ct
			minResult = v.R - v.L
		} else if v.Ct == max && minResult > v.R-v.L {
			minResult = v.R - v.L
		}
	}

	return minResult + 1
}
