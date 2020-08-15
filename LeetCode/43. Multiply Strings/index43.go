package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(multiply("5", "12"))     //562704
	fmt.Println(multiply("1234", "456")) //562704
	fmt.Println(multiply("123", "456"))  //56088
	fmt.Println(multiply("2", "3"))      //6
	fmt.Println(multiply("0", "0"))
	fmt.Println(multiply("0", "1287")) // 0
	fmt.Println(multiply("9", "9"))    // 0
}

func multiply(num1 string, num2 string) string {
	if len(num1) > len(num2) {
		tmp := num1
		num1 = num2
		num2 = tmp
	}

	var reslut []byte
	for i := len(num1) - 1; i >= 0; i-- {
		reslut = append(reslut, num1[i])
	}
	num1 = string(reslut)
	reslut = []byte{}
	for i := len(num2) - 1; i >= 0; i-- {
		reslut = append(reslut, num2[i])
	}
	num2 = string(reslut)

	result := make([]int, len(num1)*len(num2)+1)

	for i, val1 := range num1 {
		// 用num1 的每一位去乘以num2
		for j, val2 := range num2 {
			iv1, _ := strconv.Atoi(string(val1))
			iv2, _ := strconv.Atoi(string(val2))
			temp := iv1 * iv2
			result[i+j] += temp
		}
	}

	resultStr := ""
	for i, val := range result {
		if val >= 10 {
			result[i] = val % 10
			result[i+1] += val / 10
		}
	}

	flag := true
	for i := len(result) - 1; i >= 0; i-- {
		if flag && result[i] == 0 {

		} else if !flag {
			resultStr += strconv.Itoa(result[i])
		} else {
			resultStr += strconv.Itoa(result[i])
			flag = false
		}

	}
	if resultStr == "" {
		resultStr = "0"
	}
	return resultStr
}
