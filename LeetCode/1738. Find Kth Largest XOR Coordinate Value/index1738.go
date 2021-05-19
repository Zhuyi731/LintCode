package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(kthLargestValue([][]int{{9}}, 1))
	fmt.Println(kthLargestValue([][]int{{5, 2}, {1, 6}}, 1))
	fmt.Println(kthLargestValue([][]int{{5, 2}, {1, 6}}, 2))
	fmt.Println(kthLargestValue([][]int{{5, 2}, {1, 6}}, 3))
	fmt.Println(kthLargestValue([][]int{{5, 2}, {1, 6}}, 4))
	fmt.Println(kthLargestValue([][]int{{10, 9, 5}, {2, 0, 4}, {1, 0, 9}, {3, 4, 8}}, 10))
}

type KthList struct {
	Val    []int
	Length int
	K      int
}

func (list *KthList) Add(val int) {
	BinaryInsert(&list.Val, val)
	if list.Length == list.K {
		list.Val = list.Val[:len(list.Val)-1]
	} else {
		list.Length++
	}
}

func BinaryInsert(arr *[]int, val int) {
	l := 0
	arrLen := len(*arr)
	r := arrLen
	if arrLen == 0 {
		(*arr) = append((*arr), val)
		return
	}
	var index int
	for {
		mid := (l + r) / 2
		if val > (*arr)[mid] {
			l = mid + 1
		} else {
			r = mid
		}
		if l >= r {
			index = r
			break
		}
	}

	if index > arrLen {
		(*arr) = append((*arr), val)
	} else if index == 0 {
		(*arr) = append([]int{val}, (*arr)...)
	} else {
		tmp := append([]int{}, (*arr)[index:]...)
		slice := append((*arr)[:index], val)
		(*arr) = append(slice, tmp...)
	}

}

func kthLargestValue(matrix [][]int, k int) int {
	sum := make([][]int, len(matrix))
	total := []int{}

	for i := 0; i < len(matrix); i++ {
		sum[i] = make([]int, len(matrix[i]))
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if j == 0 {
				if i == 0 {
					sum[i][j] = matrix[i][j]
				} else {
					sum[i][j] = sum[i-1][j] ^ matrix[i][j]
				}
			} else {
				if i == 0 {
					sum[i][j] = sum[i][j-1] ^ matrix[i][j]
				} else {
					sum[i][j] = sum[i-1][j-1] ^ sum[i-1][j] ^ sum[i][j-1] ^ matrix[i][j]
				}
			}
			total = append(total, sum[i][j])
		}
	}

	sort.Ints(total)

	return total[len(matrix)*len(matrix[0])-k]
}
