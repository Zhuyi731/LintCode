package main

import (
	"fmt"
)

func main() {
	fmt.Println(topKFrequent([]string{"the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"}, 4))
	fmt.Println(topKFrequent([]string{"i", "love", "leetcode", "i", "love", "coding"}, 2))
}

func BinaryInsert(arr *[]interface{}, val interface{}, compareFunc func(a, b interface{}) int) {
	l := 0
	arrLen := len(*arr)
	r := arrLen
	if arrLen == 0 {
		(*arr) = append((*arr), val)
		return
	}
	var index int
	for {
		mid := (l + r) / 2
		if compareFunc(val, (*arr)[mid]) == 1 {
			l = mid + 1
		} else {
			r = mid
		}
		if l >= r {
			index = r
			break
		}
	}

	if index > arrLen {
		(*arr) = append((*arr), val)
	} else if index == 0 {
		(*arr) = append([]interface{}{val}, (*arr)...)
	} else {
		tmp := append([]interface{}{}, (*arr)[index:]...)
		slice := append((*arr)[:index], val)
		(*arr) = append(slice, tmp...)
	}
}

type record struct {
	Val int
	Str string
}

func topKFrequent(words []string, k int) []string {
	hashMap := map[string]int{}
	for _, v := range words {
		hashMap[v]++
	}

	kList := []interface{}{}
	for k, v := range hashMap {
		BinaryInsert(&kList,
			record{
				Val: v,
				Str: k,
			},
			func(a interface{}, b interface{}) int {
				Ra := a.(record)
				Rb := b.(record)
				if Ra.Val > Rb.Val {
					return -1
				} else if Ra.Val == Rb.Val {
					if Ra.Str < Rb.Str {
						return -1
					}
					return 1
				}
				return 1
			})
	}

	kList = kList[:k]
	result := []string{}
	for _, v := range kList {
		r := v.(record)
		result = append(result, r.Str)
	}

	return result
}
