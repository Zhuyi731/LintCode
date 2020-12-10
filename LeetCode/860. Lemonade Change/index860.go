package main

import "fmt"

func main() {
	fmt.Println(lemonadeChange([]int{5, 5, 10, 20, 5, 5, 5, 5, 5, 5, 5, 5, 5, 10, 5, 5, 20, 5, 20, 5}))
	fmt.Println(lemonadeChange([]int{5, 5, 5, 10, 20}))
	fmt.Println(lemonadeChange([]int{5, 5, 10}))
	fmt.Println(lemonadeChange([]int{5, 5, 10, 10, 20}))
	fmt.Println(lemonadeChange([]int{10, 10, 20}))
}

func lemonadeChange(bills []int) bool {
	currentBills := []int{0, 0}
	for i, v := range bills {
		if v == 5 {
			currentBills[0]++
		} else if v == 10 {
			if currentBills[0] > 0 {
				currentBills[0]--
				currentBills[1]++
			} else {
				fmt.Println(i)
				return false
			}
		} else {
			if currentBills[0] > 0 && currentBills[1] > 0 {
				currentBills[0]--
				currentBills[1]--
			} else if currentBills[0] > 2 {
				currentBills[0] = currentBills[0] - 3
			} else {
				fmt.Println(i)
				return false
			}
		}
	}
	return true
}
