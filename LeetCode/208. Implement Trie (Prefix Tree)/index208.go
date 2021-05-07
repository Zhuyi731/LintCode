package main

import (
	"fmt"
)

func main() {
	obj := Constructor()
	obj.Insert("apple")
	param_2 := obj.Search("apple")
	param_3 := obj.Search("app")
	param_4 := obj.StartsWith("apple")
	param_5 := obj.StartsWith("app")
	obj.Insert("app")
	param_6 := obj.Search("app")
	// 1 0 1 1
	fmt.Println(param_2, param_3, param_4, param_5, param_6)
}

type Trie struct {
	Nodes map[rune]*Trie
	End   bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{
		Nodes: make(map[rune]*Trie),
		End:   false,
	}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	for _, v := range word {
		if this.Nodes[v] == nil {
			t := Constructor()
			this.Nodes[v] = &t
		}
		this = this.Nodes[v]
	}
	this.End = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	for _, v := range word {
		if this.Nodes[v] != nil {
			this = this.Nodes[v]
		} else {
			return false
		}
	}
	// 如果不是叶子节点也要返回false
	if this.End {
		return true
	}
	return false
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	for _, v := range prefix {
		if this.Nodes[v] != nil {
			this = this.Nodes[v]
		} else {
			return false
		}
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
