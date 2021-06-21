package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(readBinaryWatch((8)))
}

func fillZero(h, m int) string {
	var sh, sm string

	sh = strconv.Itoa(h)
	if m < 10 {
		sm = "0" + strconv.Itoa(m)
	} else {
		sm = strconv.Itoa(m)
	}
	return sh + ":" + sm
}

func readBinaryWatch(turnedOn int) []string {
	if turnedOn > 8 {
		return []string{}
	}
	a := []int{1, 2, 4, 8}
	b := []int{1, 2, 4, 8, 16, 32}

	result := []string{}
	var dfs func(selectedA, selectedB, curA, curB, curUsed int)
	dfs = func(selectedA, selectedB, curA, curB, curUsed int) {
		if curA >= 12 || curB >= 60 || (selectedB >= 6 && curUsed < turnedOn) {
			return
		}
		if curUsed == turnedOn {
			result = append(result, fillZero(curA, curB))
			return
		}
		if selectedA == 4 {
			// 开始选B
			dfs(selectedA, selectedB+1, curA, curB+b[selectedB], curUsed+1)
			dfs(selectedA, selectedB+1, curA, curB, curUsed)
		} else {
			//否则继续选A
			// 不选
			dfs(selectedA+1, selectedB, curA+a[selectedA], curB, curUsed+1)
			dfs(selectedA+1, selectedB, curA, curB, curUsed)
		}
	}
	dfs(0, 0, 0, 0, 0)
	return result
}
