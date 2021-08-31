package main

import "fmt"

func main() {
	fmt.Println(corpFlightBookings([][]int{{1, 2, 10}, {2, 3, 20}, {2, 5, 25}}, 5))
}

func corpFlightBookings(bookings [][]int, n int) []int {
	result := make([]int, n)
	for _, v := range bookings {
		result[v[0]-1] += v[2]
		if v[1] < n {
			result[v[1]] -= v[2]
		}
	}

	for i := 1; i < len(result); i++ {
		result[i] += result[i-1]
	}

	return result
}
