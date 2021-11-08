package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(getHint("1807", "7810"))
	fmt.Println(getHint("1123", "0111"))
	fmt.Println(getHint("1", "0"))
	fmt.Println(getHint("1", "1"))
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func getHint(secret string, guess string) string {
	mp1, mp2 := map[rune]int{}, map[rune]int{}
	bulls := 0
	cows := 0
	for i, v := range secret {
		if v == rune(guess[i]) {
			bulls++
		} else {
			mp1[v]++
			mp2[rune(guess[i])]++
		}
	}
	for k, v := range mp1 {
		cows += min(v, mp2[k])
	}
	return strconv.Itoa(bulls) + "A" + strconv.Itoa(cows) + "B"
}
