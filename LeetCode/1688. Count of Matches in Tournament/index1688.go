package main

import "fmt"

func main() {
	fmt.Println(numberOfMatches(7))
	fmt.Println(numberOfMatches(14))
	fmt.Println(numberOfMatches(140))
	fmt.Println(numberOfMatches(141))
}

// 每次只能淘汰一个人  所以为n-1
func numberOfMatches(n int) int {
	return n - 1
}
