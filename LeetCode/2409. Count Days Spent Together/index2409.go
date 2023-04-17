package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(countDaysTogether("01-15", "02-18", "01-16", "02-19"))
	fmt.Println(countDaysTogether("08-15", "08-18", "08-16", "08-19"))
	fmt.Println(countDaysTogether("08-15", "08-23", "08-16", "08-19"))
	fmt.Println(countDaysTogether("08-17", "08-24", "08-16", "08-19"))
	fmt.Println(countDaysTogether("08-17", "08-22", "08-16", "08-30"))
	fmt.Println(countDaysTogether("10-01", "10-31", "11-01", "12-31"))
}

var mDays = []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

func getDay(date string) int {
	d := strings.Split(date, "-")
	month, _ := strconv.Atoi(d[0])
	day, _ := strconv.Atoi(d[1])
	result := day
	for i := 0; i < month-1; i++ {
		result += mDays[i]
	}
	return result
}

func countDaysTogether(arriveAlice string, leaveAlice string, arriveBob string, leaveBob string) int {
	if leaveAlice < arriveBob || arriveAlice > leaveBob { // without touching
		return 0
	}

	aa := getDay(arriveAlice)
	la := getDay(leaveAlice)
	ab := getDay(arriveBob)
	lb := getDay(leaveBob)

	left := aa
	if ab > aa {
		left = ab
	}
	right := la
	if lb < la {
		right = lb
	}

	return right - left + 1
}
