package main

import "fmt"

func main() {
	fmt.Println(validMountainArray([]int{14, 82, 89, 84, 79, 70, 70, 68, 67, 66, 63, 60, 58, 54, 44, 43, 32, 28, 26, 25, 22, 15, 13, 12, 10, 8, 7, 5, 4, 3}))
	fmt.Println(validMountainArray([]int{0, 1, 2, 4, 2, 1}))
	fmt.Println(validMountainArray([]int{2, 1, 2, 3, 5, 7, 9, 10, 12, 14, 15, 16, 18, 14, 13}))
	fmt.Println(validMountainArray([]int{5, 4, 3, 2}))
	fmt.Println(validMountainArray([]int{1, 3, 2}))
	fmt.Println(validMountainArray([]int{2, 1}))
	fmt.Println(validMountainArray([]int{3, 5, 5}))
	fmt.Println(validMountainArray([]int{0, 3, 2, 1}))
}
func validMountainArray(A []int) bool {
	if len(A) < 3 {
		return false
	}

	findUp := false
	findTop := false
	for i := 1; i < len(A); i++ {
		if !findTop {
			if A[i] < A[i-1] {
				if findUp {
					findTop = true
				} else {
					return false
				}
			} else if !findUp && A[i] > A[i-1] {
				findUp = true
			} else if A[i] == A[i-1] {
				return false
			}
		} else {
			if A[i] >= A[i-1] {
				return false
			}
		}
	}
	return findTop
}
