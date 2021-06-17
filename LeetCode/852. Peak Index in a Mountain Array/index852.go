package main

import "fmt"

func main() {
	fmt.Println(peakIndexInMountainArray([]int{0, 1, 0}))
	fmt.Println(peakIndexInMountainArray([]int{0, 2, 1, 0}))
	fmt.Println(peakIndexInMountainArray([]int{0, 10, 5, 2}))
	fmt.Println(peakIndexInMountainArray([]int{0, 4, 5, 2}))
	fmt.Println(peakIndexInMountainArray([]int{24, 69, 100, 99, 79, 78, 67, 36, 26, 19})) // 2
}

func peakIndexInMountainArray(arr []int) int {
	l := 0
	r := len(arr) - 1
	m := (l + r) / 2
	checkOk := func() int {
		if arr[m] > arr[m-1] && arr[m] > arr[m+1] {
			return 0
		} else if arr[m] < arr[m-1] {
			return -1
		} else {
			return 1
		}
	}
	for {
		ok := checkOk()
		if ok == 0 {
			return m
		} else if ok == -1 {
			r = m
		} else {
			l = m + 1
		}
		m = (l + r) / 2
	}
}
