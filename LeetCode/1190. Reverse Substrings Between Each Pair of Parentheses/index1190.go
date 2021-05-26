package main

import (
	"fmt"
)

func main() {
	fmt.Println(reverseParentheses("(u(love)i)"))
	fmt.Println(reverseParentheses("(abcd)"))
	fmt.Println(reverseParentheses("(ed(et(oc))el)"))
	fmt.Println(reverseParentheses("a(bcdefghijkl(mno)p)q"))
}

type stack struct {
	Val []rune
}

func (s *stack) Push(r rune) {
	s.Val = append(s.Val, r)
}

func (s *stack) PushStr(str string) {
	s.Val = append(s.Val, []rune(str)...)
}

func (s *stack) Pop() string {
	l := len(s.Val)
	r := make([]rune, l)
	for i, v := range s.Val {
		r[l-i-1] = v
	}
	return string(r)
}

func reverseParentheses(s string) string {
	result := ""
	shouldPush := false
	st := []stack{}
	for _, v := range s {
		if v == '(' {
			shouldPush = true
			st = append(st, stack{})
		} else if v == ')' {
			tmp := st[len(st)-1].Pop()
			st = st[:len(st)-1]
			if len(st) == 0 {
				shouldPush = false
				result += tmp
			} else {
				st[len(st)-1].PushStr(tmp)
			}
		} else {
			if shouldPush {
				st[len(st)-1].Push(v)
			} else {
				result += string(v)
			}
		}
	}
	return result
}
