package main

import "fmt"

func main() {
	fmt.Println(truncateSentence("Hello how are you Contestant", 4))
	fmt.Println(truncateSentence("What is the solution to this problem", 4))
	fmt.Println(truncateSentence("chopper is not a tanuki", 5))
}

func truncateSentence(s string, k int) string {
	index := len(s)
	for i, v := range s {
		if v == ' ' {
			k--
			if k == 0 {
				index = i
				break
			}
		}
	}
	return s[:index]
}
