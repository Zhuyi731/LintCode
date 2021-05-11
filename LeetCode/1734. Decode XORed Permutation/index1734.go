package main

import "fmt"

func main() {
	fmt.Println(decode([]int{3, 1}))
	fmt.Println(decode([]int{6, 5, 4, 6}))
}

func decode(encoded []int) []int {
	l := len(encoded)
	all := 1
	for i := 2; i <= l+1; i++ {
		all = i ^ all
	}
	v1 := 0
	for i := 1; i < l; i = i + 2 {
		v1 = v1 ^ encoded[i]
	}
	origin := make([]int, l+1)
	origin[0] = v1 ^ all
	for i := 1; i <= l; i++ {
		origin[i] = encoded[i-1] ^ origin[i-1]
	}
	return origin
}
