package main

import "fmt"

func main() {
	fmt.Println(findComplement(2147483646))
	fmt.Println(findComplement(5))
	fmt.Println(findComplement(2))
	fmt.Println(findComplement(1))
}

var l []int

func init() {
	cur := 1
	l = []int{-1}
	for i := 1; i <= 31; i++ {
		l = append(l, (cur<<i)-1)
	}
}

func findComplement(num int) int {
	for i := 1; i <= 31; i++ {
		if num > l[i-1] && num <= l[i] {
			return num ^ l[i]
		}
	}
	return 0
}
