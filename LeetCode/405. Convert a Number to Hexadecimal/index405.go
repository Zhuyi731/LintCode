package main

import "fmt"

func main() {
	toHex(26)
	toHex(-1)
}

var rest = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f"}
var max = 0xffffffff

func toHex(num int) string {
	if num == 0 {
		return "0"
	}
	if num < 0 {
		num = max + num + 1
	}
	result := ""
	for num != 0 {
		r := num % 16
		result = rest[r] + result
		num = num / 16
	}
	fmt.Println(result)

	return result
}
