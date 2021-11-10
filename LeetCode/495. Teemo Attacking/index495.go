package main

import "fmt"

func main() {
	fmt.Println(findPoisonedDuration([]int{1, 4}, 2))                      //4
	fmt.Println(findPoisonedDuration([]int{1, 2}, 2))                      //3
	fmt.Println(findPoisonedDuration([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 1)) //3
}

func findPoisonedDuration(timeSeries []int, duration int) int {
	total := 0
	for k, v := range timeSeries {
		if k == len(timeSeries)-1 || timeSeries[k+1]-v >= duration {
			total += duration
		} else {
			total += timeSeries[k+1] - v
		}
	}
	return total
}
