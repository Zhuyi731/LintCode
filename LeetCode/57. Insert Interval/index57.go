package main

import "fmt"

func main() {
	fmt.Println(insert([][]int{{3, 5}, {12, 15}}, []int{6, 6})) // 15 69
	fmt.Println(insert([][]int{{1, 3}, {6, 9}}, []int{2, 5}))   // 15 69
	fmt.Println(insert([][]int{{1, 5}}, []int{0, 0}))           // 00 15
	fmt.Println(insert([][]int{{1, 5}}, []int{0, 3}))           // 0 5
	fmt.Println(insert([][]int{{1, 5}}, []int{6, 7}))           // 1 5 6  7
	fmt.Println(insert([][]int{{1, 15}, {2, 4}}, []int{3, 16})) // 1 16
	fmt.Println(insert([][]int{{1, 3}, {6, 9}}, []int{2, 5}))   // 1 5 6 9
	fmt.Println(insert([][]int{}, []int{5, 7}))                 // 5 7
	fmt.Println(insert([][]int{{1, 5}}, []int{2, 3}))           // 1 5
	fmt.Println(insert([][]int{{1, 5}}, []int{2, 7}))           // 1 7
	fmt.Println(insert([][]int{{1, 15}, {2, 4}}, []int{3, 6}))  // 1 15
}

func insert(intervals [][]int, newInterval []int) [][]int {
	l := len(intervals)
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}
	if newInterval[1] < intervals[0][0] {
		return append([][]int{newInterval}, intervals...)
	}
	insertFlag := -1
	i := 0
	endFlag := -1

	maxR := newInterval[1]
	for ; i < l; i++ {
		arr := intervals[i]
		if insertFlag == -1 && arr[1] >= newInterval[0] { // 需要插入了
			insertFlag = i
		}

		if insertFlag != -1 && arr[0] > newInterval[1] {
			// 开始合并
			endFlag = i
			break
		}

		if insertFlag != -1 {
			if arr[1] > maxR {
				maxR = arr[1]
			}
		}
	}
	var mergedInterval []int
	if endFlag == -1 {
		endFlag = l
	} else {
		if intervals[endFlag-1][1] > maxR {
			maxR = intervals[endFlag-1][1]
		}
	}
	if insertFlag == -1 {
		insertFlag = l
		mergedInterval = newInterval
	} else if insertFlag == endFlag {
		mergedInterval = newInterval
	} else {
		minL := intervals[insertFlag][0]
		if newInterval[0] < minL {
			minL = newInterval[0]
		}
		mergedInterval = []int{minL, maxR}
	}
	// tmp := intervals[endFlag:]
	intervals = append(intervals[:insertFlag], append([][]int{mergedInterval}, intervals[endFlag:]...)...)
	// intervals = append(intervals, tmp...)
	return intervals
}
