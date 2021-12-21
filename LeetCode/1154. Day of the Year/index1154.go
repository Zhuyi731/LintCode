package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(dayOfYear("2019-01-09")) // 9
	fmt.Println(dayOfYear("2019-02-10")) // 41
	fmt.Println(dayOfYear("2003-03-01")) //  60
	fmt.Println(dayOfYear("2004-03-01")) // 61
	fmt.Println(dayOfYear("2000-03-01")) // 61
	fmt.Println(dayOfYear("1900-03-01")) // 61
}

func dayOfYear(date string) int {
	arr := strings.Split(date, "-")
	y, m, d := arr[0], arr[1], arr[2]
	year, _ := strconv.Atoi(y)
	month, _ := strconv.Atoi(m)
	day, _ := strconv.Atoi(d)
	days := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	isLunaYear := func(y int) bool {
		return y%4 == 0 && (y%100 != 0 || y%400 == 0)
	}
	if isLunaYear(year) {
		days[1]++
	}
	result := 0
	for i := 0; i < month-1; i++ {
		result += days[i]
	}

	result += day
	return result
}
