package main

import "fmt"

func main() {
	action( // [null,null,null,null,null,false,false,null,true,true,false,false,true,false]
		[]string{"WordDictionary", "addWord", "addWord", "addWord", "addWord", "search", "search", "addWord", "search", "search", "search", "search", "search", "search"},
		[]string{"", "at", "and", "an", "add", "a", ".at", "bat", ".at", "an.", "a.d.", "b.", "a.d", "."},
	)
	action( // true,true,false,true,false,false
		[]string{"WordDictionary", "addWord", "addWord", "search", "search", "search", "search", "search", "search"},
		[]string{"", "a", "a", ".", "a", "aa", "a", ".a", "a."},
	)
	action(
		[]string{"WordDictionary", "addWord", "addWord", "addWord", "search", "search", "search", "search"},
		[]string{"", "bad", "dad", "mad", "pad", "bad", ".ad", "b.."},
	)
}

func action(act []string, params []string) {
	var wd WordDictionary
	for i, v := range act {
		switch v {
		case "WordDictionary":
			wd = Constructor()
			fmt.Println("\n\n\n\n")
			fmt.Println("null")
		case "addWord":
			wd.AddWord(params[i])
			fmt.Println("null")
		case "search":
			fmt.Println(wd.Search(params[i]))
		}
	}

}

type WordDictionary struct {
	ChildNodes map[rune]*WordDictionary
	IsEndPoint bool
}

func Constructor() WordDictionary {
	return WordDictionary{
		ChildNodes: map[rune]*WordDictionary{},
		IsEndPoint: false,
	}
}

func (this *WordDictionary) AddWord(word string) {
	cur := this
	var t *WordDictionary
	for _, v := range word {
		if _, ok := cur.ChildNodes[v]; !ok {
			wd := Constructor()
			cur.ChildNodes[v] = &wd
		}
		t = cur.ChildNodes[v]
		cur = t
	}
	t.IsEndPoint = true
}

func (this *WordDictionary) Search(word string) bool {
	if word == "" {
		return this.IsEndPoint
	}
	if len(this.ChildNodes) == 0 {
		return false
	}
	for i, v := range word {
		if v == '.' {
			newWord := word[i+1:]
			for vv := range this.ChildNodes {
				t := this.ChildNodes[vv]
				if t.Search(newWord) == true {
					return true
				}
			}
			return false
		} else if _, ok := this.ChildNodes[v]; ok {
			t := this.ChildNodes[v]
			this = t
		} else {
			return false
		}
	}
	return this.IsEndPoint
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
