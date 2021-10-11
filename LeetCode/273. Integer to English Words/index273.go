package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println(numberToWords(100000))
	fmt.Println(numberToWords(1000))
	fmt.Println(numberToWords(100))
	fmt.Println(numberToWords(1100))
	fmt.Println(numberToWords(200))
	fmt.Println(numberToWords(2000))
	fmt.Println(numberToWords(1000001))
	fmt.Println(numberToWords(999999))
	fmt.Println(numberToWords(1000000))
	fmt.Println(numberToWords(1000))
	fmt.Println(numberToWords(20))
	fmt.Println(numberToWords(1))
	fmt.Println(numberToWords(19))
	fmt.Println(numberToWords(123))        // One Hundred Twenty Three
	fmt.Println(numberToWords(12345))      // Twelve Thousand Three Hundred Forty Five
	fmt.Println(numberToWords(1234567))    // One Million Two Hundred Thirty Four Thousand Five Hundred Sixty Seven
	fmt.Println(numberToWords(1234567891)) // One Billion Two Hundred Thirty Four Million Five Hundred Sixty Seven Thousand Eight Hundred Ninety One
}

var units = []string{"Thousand", "Million", "Billion"}
var bigTen = []string{"", "", "Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety"}
var normal = []string{"", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen"}

func numberToWords(num int) string {
	if num == 0 {
		return "Zero"
	}
	reg := regexp.MustCompile("(^\\s*|\\s*$)")
	val, _ := rec(num, 0)
	result := reg.ReplaceAllString(val, "")
	reg2 := regexp.MustCompile(("\\s{2,}"))
	result2 := reg2.ReplaceAllString(result, " ")
	return result2
}

func rec(num int, depth int) (string, bool) {
	r := num % 1000
	result := ""
	if r >= 100 {
		hundredNum := r / 100
		result += normal[hundredNum] + " Hundred"
		r = r % 100
	}
	if r >= 20 {
		tenNum := r / 10
		result = result + " " + bigTen[tenNum]
		result = result + " " + normal[r%10]
	} else if r == 0 {
		result += ""
	} else {
		result = result + " " + normal[r]
	}

	if num >= 1000 {
		//遍历
		count, needUnit := rec(num/1000, depth+1)
		if needUnit {
			result = count + " " + units[depth] + " " + result
		} else {
			result = count + " " + result
		}
	}
	if num%1000 == 0 {
		return result, false
	}
	return result, true
}
