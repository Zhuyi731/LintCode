package main

import "fmt"

func main() {
	fmt.Println(luckyNumbers([][]int{{3, 7, 8}, {9, 11, 13}, {15, 16, 17}}))          //15
	fmt.Println(luckyNumbers([][]int{{1, 10, 4, 2}, {9, 3, 8, 7}, {15, 16, 17, 12}})) //12
	fmt.Println(luckyNumbers([][]int{{7, 8}, {1, 2}}))                                //7
}

func luckyNumbers(matrix [][]int) []int {
	result := []int{}
	for _, row := range matrix {
		currentMin := row[0]
		currentMinIndex := 0
		for j, v := range row {
			if v < currentMin {
				currentMin = v
				currentMinIndex = j
			}
		}
		isOk := true
		for t := 0; t < len(matrix); t++ {
			if matrix[t][currentMinIndex] > currentMin {
				isOk = false
				break
			}
		}
		if isOk {
			result = append(result, currentMin)
		}
	}
	return result
}
