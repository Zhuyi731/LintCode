package main

import (
	"fmt"
)

func main() {
	fmt.Println(sortString("acbzxhuiebzxhncioua"))
	fmt.Println(sortString("rat"))
	fmt.Println(sortString("aaaabbbbcccc"))
}

func sortString(s string) string {
	hashMap := ['z' + 1]int{}

	for _, ch := range s {
		hashMap[ch]++
	}

	n := len(s)
	result := make([]byte, 0, n)
	for {
		if len(result) == n {
			break
		}
		for i := byte('a'); i <= 'z'; i++ {
			if hashMap[i] != 0 {
				hashMap[i] = hashMap[i] - 1
				result = append(result, i)
			}
		}

		for i := byte('z'); i >= 'a'; i-- {
			if hashMap[i] != 0 {
				hashMap[i] = hashMap[i] - 1
				result = append(result, i)
			}
		}
	}
	return string(result)
}
