package main

import "fmt"

func main() {
	fmt.Println(findMaxLength([]int{0, 1}))
	fmt.Println(findMaxLength([]int{0, 1, 0}))
	fmt.Println(findMaxLength([]int{1, 1, 1, 0, 1, 0, 0, 1, 0}))
}
func findMaxLength(nums []int) int {
	m := map[int]int{
		0: -1,
	}
	result := 0
	currentDis := 0
	for i, v := range nums {
		if v == 0 {
			currentDis--
		} else {
			currentDis++
		}
		if index, ok := m[currentDis]; ok {
			if result < i-index {
				result = i - index
			}
		} else {
			m[currentDis] = i
		}
	}
	return result
}
