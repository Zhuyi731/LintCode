package main

import "fmt"

func main() {
	fmt.Println(partitionLabels("ababcbacadefegdehijhklij"))
}

func Max(num ...int) int {
	var maxNum int
	for index, val := range num {
		if index == 0 {
			maxNum = val
			continue
		} else {
			if val > maxNum {
				maxNum = val
			}
		}
	}
	return maxNum
}

func partitionLabels(S string) []int {
	lastIndexList := make(map[rune]int)

	for index, value := range S {
		lastIndexList[value] = index
	}

	result := make([]int, 0)

	start := 0 //ababcbacadefegdehijhklij
	end := 0
	for index, value := range S {
		end = Max(end, lastIndexList[value]) // 拓展end
		if end == index {
			// 说明区间到头了
			currentResult := end - start + 1
			start = end + 1
			end = start
			result = append(result, currentResult)
		}
	}
	return result
}
