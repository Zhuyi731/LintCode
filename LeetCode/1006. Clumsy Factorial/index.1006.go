package main

import "fmt"

func main() {
	fmt.Println(clumsy(10000))
	fmt.Println(clumsy(4))
	fmt.Println(clumsy(10))
	fmt.Println(clumsy(9999))
}

func clumsy(n int) int {
	tmp := []int{0, 1, 2, 6, 7}
	if n <= 4 {
		return tmp[n]
	} else {
		tmp := []int{5, 1, 2, 6}
		result := n*(n-1)/(n-2) + n - 3
		for i := n - 4; i > 4; i = i - 4 {
			result = result - int(i*(i-1)/(i-2)) + i - 3
		}
		return result - (tmp[n%4])
	}
}
