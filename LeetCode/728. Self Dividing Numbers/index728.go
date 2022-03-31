package main

import "fmt"

func main() {
	fmt.Println(selfDividingNumbers(1, 22))
	fmt.Println(selfDividingNumbers(47, 85))
}

var result []int

func init() {
	for i := 1; i <= 11111; i++ {
		cur := i
		for {
			if cur == 0 {
				result = append(result, i)
				break
			}
			tail := cur % 10
			cur = cur / 10
			if tail == 0 || i%tail != 0 {
				break
			}
		}

	}
}

func selfDividingNumbers(left int, right int) []int {
	l, r := -1, -1
	for i := 0; i < len(result); i++ {
		if result[i] >= left && l == -1 {
			l = i
		}
		if result[i] > right && r == -1 {
			r = i
		}
	}
	return result[l:r]
}
