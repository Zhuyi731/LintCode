package main

import "fmt"

func main() {
	fmt.Println(intToRoman(3))
	fmt.Println(intToRoman(4))
	fmt.Println(intToRoman(58))
	fmt.Println(intToRoman(1998))
	fmt.Println(intToRoman(1994))
}
func intToRoman(num int) string {
	list := []string{"I", "V", "X", "L", "C", "D", "M"}
	result := ""
	ct := 0

	fill := func(n int, chr string) string {
		full := ""
		for i := 0; i < n; i++ {
			full += chr
		}
		return full
	}
	for {
		if num == 0 {
			break
		}
		tail := num % 10
		if tail == 4 || tail == 9 {
			//       ""   +  单位  + 次级单位
			result = list[ct*2] + list[ct*2+1+tail/5] + result
		} else {
			result = fill(tail%5, list[ct*2]) + result
			if tail >= 5 {
				result = list[ct*2+1] + result
			}
		}
		num = num / 10
		ct++
	}
	return result
}
