package main

import "fmt"

func main() {
	fmt.Println(hIndex([]int{0, 1, 3, 5, 6}))
}

func hIndex(citations []int) int {
	l := len(citations)
	for i, v := range citations {
		if v >= l-i {
			return l - i
		}
	}
	return -1
}
