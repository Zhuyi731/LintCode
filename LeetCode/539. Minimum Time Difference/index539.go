package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(findMinDifference([]string{"23:59", "00:00"}))
	fmt.Println(findMinDifference([]string{"23:59", "00:00", "23:59"}))
	fmt.Println(findMinDifference([]string{"21:59", "00:00", "20:59"}))
}

func findMinDifference(timePoints []string) int {
	result := 24 * 60
	if len(timePoints) >= result {
		return 0
	}
	intTime := make([]int, len(timePoints))
	for i, v := range timePoints {
		arr := strings.Split(v, ":")
		h, m := arr[0], arr[1]
		hi, _ := strconv.ParseInt(h, 10, 32)
		mi, _ := strconv.ParseInt(m, 10, 32)
		intTime[i] = int(hi*60 + mi)
	}
	sort.Slice(intTime, func(i, j int) bool {
		return intTime[i] < intTime[j]
	})

	for i := 1; i < len(intTime); i++ {
		if intTime[i]-intTime[i-1] < result {
			result = intTime[i] - intTime[i-1]
		}
	}
	if intTime[0]+24*60-intTime[len(intTime)-1] < result {
		result = intTime[0] + 24*60 - intTime[len(intTime)-1]
	}
	return result
}
