package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Join(fullJustify([]string{
		"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do",
	}, 20), "\r\n"))
	fmt.Println(strings.Join(fullJustify([]string{"What", "must", "be", "acknowledgment", "shall", "be"}, 16), "\r\n"))
	fmt.Println(strings.Join(fullJustify([]string{"This", "is", "an", "example", "of", "text", "justification."}, 16), "\r\n"))
}

func handle(result *[]string, tmp []string, maxWidth int) {
	lt := len(tmp)
	l := len(strings.Join(tmp, ""))

	var space, spaceRest int
	if lt == 1 {
		space = (maxWidth - l)
		spaceRest = 0
	} else {
		space = (maxWidth - l) / (lt - 1)
		spaceRest = (maxWidth - l) % (lt - 1) // 不整除时多出来的
	}
	if len(tmp) == 1 {
		*result = append(*result, tmp[0]+strings.Repeat(" ", space))
		return
	}
	cur := ""
	for i, str := range tmp {
		if i == lt-1 {
			cur += str
		} else {
			cur += str + strings.Repeat(" ", space)
		}
		if i < spaceRest {
			cur += " "
		}
	}
	*result = append(*result, cur)
}

func handleLastLine(result *[]string, tmp []string, maxWidth int) {
	cur := strings.Join(tmp, " ")
	*result = append(*result, cur+strings.Repeat(" ", maxWidth-len(cur)))
}

func fullJustify(words []string, maxWidth int) []string {
	result := []string{}
	tmp := []string{}
	curLen := 0
	lw := len(words)
	for i := 0; i <= lw; i++ {
		if i == lw {
			handleLastLine(&result, tmp, maxWidth)
			break
		}
		v := words[i]
		lv := len(v)
		lt := len(tmp)
		if lt == 0 || (curLen+lv+1 <= maxWidth) {
			tmp = append(tmp, v)
			curLen += lv
			if len(tmp) != 1 {
				curLen++
			}
		} else {
			// 处理一行
			handle(&result, tmp, maxWidth)
			tmp = []string{v}
			curLen = lv
		}
	}

	return result
}
