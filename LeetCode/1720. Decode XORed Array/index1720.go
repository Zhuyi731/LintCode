package main

import (
	"fmt"
)

func main() {
	fmt.Println(decode([]int{1, 2, 3}, 1))
	fmt.Println(decode([]int{6, 2, 7, 3}, 4))
}

func decode(encoded []int, first int) []int {
	l := len(encoded)
	result := make([]int, l+1)
	result[0] = first
	for i, v := range encoded {
		result[i+1] = v ^ result[i]
	}
	return result
}
