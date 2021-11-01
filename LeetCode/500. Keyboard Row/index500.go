package main

import "fmt"

func main() {
	fmt.Println(findWords([]string{"Hello", "Alaska", "Dad", "Peace"}))
	fmt.Println(findWords([]string{"omk"}))
}

var c = []string{"qwertyuiop", "asdfghjkl", "zxcvbnm"}
var mp = map[rune]int{}

func init() {
	for i, v := range c {
		for _, r := range v {
			mp[r] = i
		}
	}
}

func findWords(words []string) []string {
	result := []string{}
	for _, word := range words {
		r := rune(word[0])
		if r < 97 {
			r = r + 32
		}
		i := mp[r]
		ok := true
		for _, v := range word {
			r := v
			if r < 97 {
				r = r + 32
			}
			if mp[r] != i {
				ok = false
				break
			}
		}
		if ok {
			result = append(result, word)
		}
	}
	return result
}
