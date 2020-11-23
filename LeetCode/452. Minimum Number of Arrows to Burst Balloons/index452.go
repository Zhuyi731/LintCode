package main

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(findMinArrowShots(createTwoDimensionArray("[[9,12],[1,10],[4,11],[8,12],[3,9],[6,9],[6,7]]")))
	fmt.Println(findMinArrowShots(createTwoDimensionArray("[[10,16],[2,8],[1,6],[7,12]]")))
	fmt.Println(findMinArrowShots(createTwoDimensionArray("[[1,2],[3,4],[5,6],[7,8]]")))
	fmt.Println(findMinArrowShots(createTwoDimensionArray("[[1,2],[2,3],[3,4],[4,5]]")))
	fmt.Println(findMinArrowShots(createTwoDimensionArray("[[1,2]]")))
	fmt.Println(findMinArrowShots(createTwoDimensionArray("[[2,3],[2,3]]")))
}

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

func findMinArrowShots(points [][]int) int {
	if len(points) <= 1 {
		return len(points)
	}

	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})

	ct := 1
	currentIndex := points[0][0]
	currentMinRight := points[0][1]
	for _, pt := range points {
		if pt[0] <= currentIndex {
			if pt[1] < currentMinRight {
				currentMinRight = pt[1]
			}
			continue
		} else {
			currentIndex = pt[0]

			if pt[0] > currentMinRight {
				// 说明超过了
				currentMinRight = pt[1]
				ct++
			} else if pt[1] < currentMinRight {
				currentMinRight = pt[1]
			}
		}
	}
	return ct
}
