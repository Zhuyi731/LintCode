package main

import "fmt"

func main() {
	fmt.Println(buddyStrings("ab", "ba"))
	fmt.Println(buddyStrings("ab", "ab"))
	fmt.Println(buddyStrings("aa", "aa"))
	fmt.Println(buddyStrings("aaaaaaabc", "aaaaaaacb"))
}

func buddyStrings(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}
	hm := map[rune]int{}
	hasSame := false
	diffS := []rune{}
	diffGoal := []rune{}
	for k, v := range goal {
		if rune(s[k]) != v {
			diffS = append(diffS, rune(s[k]))
			diffGoal = append(diffGoal, v)
		} else {
			hm[v]++
			if hm[v] >= 2 {
				hasSame = true
			}
		}
	}
	if len(diffS) == 1 || len(diffS) > 2 {
		return false
	}

	if len(diffS) == 0 {
		return hasSame
	} else {
		return diffGoal[0] == diffS[1] && diffGoal[1] == diffS[0]
	}

}
