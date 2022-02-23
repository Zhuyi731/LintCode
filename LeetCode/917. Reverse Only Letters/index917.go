package main

import "fmt"

func main() {
	fmt.Println(reverseOnlyLetters("ab-cd"))
	fmt.Println(reverseOnlyLetters("a-bC-dEf-ghIj"))
	fmt.Println(reverseOnlyLetters("Test1ng-Leet=code-Q!"))
}

func reverseOnlyLetters(s string) string {
	l, r := 0, len(s)-1
	newS := make([]byte, len(s))
	for {
		if l > r {
			return string(newS)
		}
		for {
			if l > r {
				return string(newS)
			}
			if !((s[l] >= 97 && s[l] <= 122) || (s[l] >= 65 && s[l] <= 90)) {
				newS[l] = s[l]
				l++
			} else {
				break
			}
		}

		for {
			if l > r {
				return string(newS)
			}
			if !((s[r] >= 97 && s[r] <= 122) || (s[r] >= 65 && s[r] <= 90)) {
				newS[r] = s[r]
				r--
			} else {
				break
			}

		}

		newS[l], newS[r] = s[r], s[l]
		l++
		r--
	}
}
