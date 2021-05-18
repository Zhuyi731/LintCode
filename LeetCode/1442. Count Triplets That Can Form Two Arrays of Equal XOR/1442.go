package main

import "fmt"

func main() {
	fmt.Println(countTriplets([]int{2, 3, 1, 6, 7}))                 // 4
	fmt.Println(countTriplets([]int{1, 1, 1, 1, 1}))                 // 10
	fmt.Println(countTriplets([]int{2, 2}))                          // 0
	fmt.Println(countTriplets([]int{1, 3, 5, 7, 9}))                 // 3
	fmt.Println(countTriplets([]int{7, 11, 12, 9, 5, 2, 7, 17, 22})) // 8
}

func countTriplets(arr []int) int {
	l := len(arr)
	if l < 2 {
		return 0
	}

	sum := make([]int, l+1)
	sum[0] = arr[0]
	for i := 1; i <= l; i++ {
		sum[i] = sum[i-1] ^ arr[i-1]
	}
	ct := 0
	for i := 0; i <= l; i++ {
		for k := i + 1; k <= l; k++ {
			if sum[k] == sum[i] {
				ct = ct + k - i - 1
			}
		}
	}
	return ct
}
