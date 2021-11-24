package main

import (
	"fmt"
	"strconv"
	"strings"
)

// w2 u4  r 3 o1  f5 x6 s7 g8
func main() {
	fmt.Println(originalDigits("nnei"))
	fmt.Println(originalDigits("owoztneoer")) //0,1,2
	fmt.Println(originalDigits("fviefuro"))   //4,5
}

func originalDigits(s string) string {
	strs := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	specials := []rune{'z', 'w', 'u', 'r', 'o', 'f', 'x', 's', 'g', 'e'}
	specialsNum := []int{0, 2, 4, 3, 1, 5, 6, 7, 8, 9}
	origin := [10]int{}
	mp := map[rune]int{}

	ok := func(index int, v rune) bool {
		for _, r := range strs[specialsNum[index]] {
			if mp[r] < mp[v] || mp[r] <= 0 {
				return false
			}
		}
		return true
	}

	for _, v := range s {
		mp[v]++
	}

	result := ""
	for i, v := range specials {
		if mp[v] != 0 {
			for {
				if mp[v] > 0 && ok(i, v) {
					origin[specialsNum[i]] += 1
					for _, r := range strs[specialsNum[i]] {
						mp[r] -= 1
					}
				} else {
					break
				}
			}
		}
	}
	for i, v := range origin {
		if v != 0 {
			result += strings.Repeat(strconv.Itoa(i), v)
		}
	}
	return result
}
