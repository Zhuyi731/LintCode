package main

import "fmt"

func main() {
	fmt.Println(isLongPressedName("aa", "aaaaa"))
	fmt.Println(isLongPressedName("alex", "aaleex"))
	fmt.Println(isLongPressedName("saeed", "ssaaedd"))
	fmt.Println(isLongPressedName("leelee", "lleeelee"))
	fmt.Println(isLongPressedName("laiden", "laiden"))
	fmt.Println(isLongPressedName("alex", "aaleex"))
}

func isLongPressedName(name string, typed string) bool {
	i := 0
	j := 0

	for j < len(typed) {
		if i < len(name) && name[i] == typed[j] {
			i++
			j++
		} else if i > 0 && name[i-1] == typed[j] {
			j++
		} else {
			return false
		}
	}

	return i == len(name)
}
