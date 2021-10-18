package main

import "fmt"

func main() {
	fmt.Println(numMovesStones(3, 5, 1))
	fmt.Println(numMovesStones(4, 5, 1))
}

func numMovesStones(a int, b int, c int) []int {
	if b < a {
		a, b = b, a
	}
	if c < a {
		a, c = c, a
	}
	if c < b {
		b, c = c, b
	}
	if b == a+1 && c == b+1 {
		return []int{0, 0}
	}
	if b == a+1 || c == b+1 || b-a == 2 || c-b == 2 {
		return []int{1, c - a - 2}
	}
	return []int{2, c - a - 2}
}
