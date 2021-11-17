package main

import "fmt"

func main() {
	fmt.Println(maxProduct([]string{"abcw", "baz", "foo", "bar", "xtfn", "abcdef"}))
	fmt.Println(maxProduct([]string{"a", "ab", "abc", "d", "cd", "bcd", "abcd"}))
	fmt.Println(maxProduct([]string{"a", "aa", "aaa", "aaaa"}))
}

func maxProduct(words []string) int {
	containsMap := map[int]bool{}

	for i, v := range words {
		for _, r := range v {
			l := int(r) - 97
			containsMap[i*100+l] = true
		}
	}
	max := 0
	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			canPlus := true
			for _, r := range words[i] {
				l := int(r) - 97
				if containsMap[j*100+l] {
					canPlus = false
					break
				}
			}
			if canPlus {
				p := len(words[i]) * len(words[j])
				if p > max {
					max = p
				}
			}
		}
	}
	return max
}
