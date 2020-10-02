package main

import "fmt"

type HashSet struct {
	set map[rune]bool
}

func New() *HashSet {
	return &HashSet{make(map[rune]bool)}
}

func (set *HashSet) Set(key rune) {
	set.set[key] = true
}

func (set *HashSet) Get(key rune) bool {
	return set.set[key]
}

func main() {
	fmt.Println(numJewelsInStones("aA", "aAAbbbb"))
}

func numJewelsInStones(J string, S string) int {
	hashSet := New()
	for _, val := range J {
		hashSet.Set(val)
	}
	ct := 0
	for _, val := range S {
		ok := hashSet.Get(val)
		if ok {
			ct++
		}
	}
	return ct
}
