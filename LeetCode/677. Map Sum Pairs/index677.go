package main

import "fmt"

func main() {
	action()
}

func action() {
	ms0 := Constructor()
	ms0.Insert("apple", 3)
	fmt.Println(ms0.Sum("ap")) // 3
	ms0.Insert("app", 2)
	ms0.Insert("apple", 2)
	fmt.Println(ms0.Sum("ap")) //4

	ms := Constructor()
	ms.Insert("apple", 3)
	fmt.Println(ms.Sum("ap"))
	ms.Insert("app", 2)
	fmt.Println(ms.Sum("ap"))

	ms2 := Constructor()
	ms2.Insert("a", 3)
	fmt.Println(ms2.Sum("ap"))
	ms2.Insert("b", 2)
	fmt.Println(ms2.Sum("a"))
}

type MapSum struct {
	Childrens map[rune]*MapSum
	Val       int
}

func Constructor() MapSum {
	return MapSum{
		Childrens: map[rune]*MapSum{},
		Val:       0,
	}
}

func (this *MapSum) Insert(key string, val int) {
	cur := this
	for _, v := range key {
		if (cur.Childrens[v]) == nil {
			mp := Constructor()
			cur.Childrens[v] = &mp
		}
		mp := cur.Childrens[v]
		cur = mp
	}
	cur.Val = val
}

func (this *MapSum) Sum(prefix string) int {
	if len(prefix) == 0 {
		sum := 0
		for _, v := range this.Childrens {
			sum += v.Sum("")
		}
		return sum + this.Val
	} else {
		for _, v := range prefix {
			if this.Childrens[v] != nil {
				this = this.Childrens[v]
			} else {
				return 0
			}
		}
		return this.Sum("")
	}
}

/**
 * Your MapSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(key,val);
 * param_2 := obj.Sum(prefix);
 */
