package main 

import (
	"fmt"
)

func main(){
	fmt.Println(strStr("aaaa","aaaaa"))
}

func strStr(haystack string, needle string) int {
    if needle == ""{
		return 0
	}

	needleLen := len(needle)
	haystackLen := len(haystack)
	for i,_ := range haystack{
		if i+needleLen >haystackLen{
			break;
		}
		var isMatch = true
		for j,_ := range needle{
			if	haystack[i+j] != needle[j]{
				isMatch = false
				break;
			}
		}
		if isMatch{
			return i
		}
	}

	return -1
}

/*
Implement strStr().

Return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.

Example 1:

Input: haystack = "hello", needle = "ll"
Output: 2
Example 2:

Input: haystack = "aaaaa", needle = "bba"
Output: -1
Clarification:

What should we return when needle is an empty string? This is a great question to ask during an interview.

For the purpose of this problem, we will return 0 when needle is an empty string. This is consistent to C's strstr() and Java's indexOf().
*/