package main

import "fmt"

func main() {
	fmt.Println(maxDepth("(1+(2*3)+((8)/4))+1")) //3
	fmt.Println(maxDepth("(1)+((2))+(((3)))"))   //3
	fmt.Println(maxDepth("1+(2*3)/(2-1)"))       //1
	fmt.Println(maxDepth("1"))                   //0
}

func maxDepth(s string) int {
	max, current := 0, 0
	for _, v := range s {
		if v == '(' {
			current++
			if current > max {
				max = current
			}
		} else if v == ')' {
			current--
		}
	}
	return max
}
