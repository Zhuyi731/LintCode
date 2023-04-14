package main

import "fmt"

func main() {
	fmt.Println(mostFrequentEven([]int{0, 1, 2, 2, 4, 4, 1}))
	fmt.Println(mostFrequentEven([]int{4, 4, 4, 9, 2, 4}))
	fmt.Println(mostFrequentEven([]int{1, 1, 1, 3, 5, 8, 6, 7}))
	fmt.Println(mostFrequentEven([]int{29, 47, 21, 41, 13, 37, 25, 7}))
}

func mostFrequentEven(nums []int) int {
	hashMap := map[int]int{}
	currentMax := [2]int{-1, 0}

	for _, v := range nums {
		if v%2 == 0 {
			hashMap[v]++
			if hashMap[v] > currentMax[1] || (hashMap[v] == currentMax[1] && v < currentMax[0]) {
				currentMax[0] = v
				currentMax[1] = hashMap[v]
			}
		}
	}

	return currentMax[0]
}
