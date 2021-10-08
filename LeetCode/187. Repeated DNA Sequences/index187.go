package main

import "fmt"

func main() {
	fmt.Println(findRepeatedDnaSequences("AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"))
	fmt.Println(findRepeatedDnaSequences("AAAAAAAAAAAAA"))
}

func findRepeatedDnaSequences(s string) []string {
	m := map[string]int{}
	result := []string{}

	for i := range s {
		if i > len(s)-10 {
			break
		}
		str := s[i : i+10]
		m[str]++
		if m[str] == 2 {
			result = append(result, str)
		}
	}
	return result
}
