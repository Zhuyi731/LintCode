package main

import "fmt"

func main() {
	fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
	fmt.Println(minWindow("a", "a"))
	fmt.Println(minWindow("a", "aa"))
}

func doseMatch(a, b map[rune]int) bool {
	if len(b) == 0 {
		return false
	}

	for k, v := range b {
		if a[k] < v {
			return false
		}
	}
	return true
}

func minWindow(s string, t string) string {
	hashMap := map[rune]int{}
	for _, v := range t {
		hashMap[v]++
	}

	currentHashMap := map[rune]int{}
	l, r := 0, 0
	currentMin := []int{0, len(s)}
	hasMatch := false

	for {
		if doseMatch(currentHashMap, hashMap) {
			hasMatch = true
			if r-l < currentMin[1]-currentMin[0] {
				currentMin = []int{l, r}
			}

			// 缩减l
			currentHashMap[rune(s[l])]--
			l++
		} else {
			if r == len(s) {
				break
			}
			currentHashMap[rune(s[r])]++
			r++
		}
	}

	if !hasMatch {
		return ""
	}

	return s[currentMin[0]:currentMin[1]]
}
