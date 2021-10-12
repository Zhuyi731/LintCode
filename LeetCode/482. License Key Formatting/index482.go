package main

import (
	"fmt"
)

func main() {
	fmt.Println(licenseKeyFormatting("5F3Z-2e-9-w", 4))
	fmt.Println(licenseKeyFormatting("2-5g-3-J", 2))
}

func licenseKeyFormatting(s string, k int) string {
	result := ""
	ct := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '-' {
			continue
		}
		b := s[i]
		if s[i] >= 97 {
			b = s[i] - 32
		}
		if ct == k {
			result = string(b) + "-" + result
			ct = 0
		} else {
			result = string(b) + result
		}
		ct++
	}
	return string(result)
}
