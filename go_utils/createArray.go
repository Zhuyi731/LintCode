package util

import (
	"regexp"
	"strconv"
	"strings"
)

//[[10,16],[2,8],[1,6],[7,12]]

func createArray(input string) []int {
	input = input[1 : len(input)-1] //strip 【】
	nums := strings.Split(input, ",")
	result := make([]int, len(nums))
	for i, v := range nums {
		result[i], _ = strconv.Atoi(v)
	}

	return result
}

func createTwoDimensionArray(input string) [][]int {
	input = input[1 : len(input)-1] //strip 【】
	input = strings.Replace(input, " ", "", 0)
	splitExp := regexp.MustCompile("],")
	oneDimension := splitExp.Split(input, -1)

	result := make([][]int, len(oneDimension))
	for i, v := range oneDimension {
		if v[len(v)-1] == ']' {
			result[i] = createArray(v)
		} else {
			result[i] = createArray(v + "]")
		}
	}
	return result
}
