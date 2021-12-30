package main

import "fmt"

func main() {
	fmt.Println(repeatedStringMatch("abcd", "cdabcdab"))
	fmt.Println(repeatedStringMatch("a", "aa"))
}

func repeatedStringMatch(a string, b string) int {
	ct := 1
	la := len(a)
	for i := range a {
		curIndex := i
		ct = 1
		ok := true
		for ib, vb := range b {
			if vb == rune(a[curIndex]) {
				curIndex++
				if curIndex >= la {
					curIndex = curIndex % la
					if ib != len(b)-1 {
						ct++
					}
				}
			} else {
				ok = false
				break
			}
		}
		if ok {
			return ct
		}
	}
	return -1
}
