package main

import "fmt"

func main() {
	fmt.Println(checkPerfectNumber(28))
	fmt.Println(checkPerfectNumber(7))
	fmt.Println(checkPerfectNumber(6))
	fmt.Println(checkPerfectNumber(16))
}

func checkPerfectNumber(num int) bool {
	if num == 1 {
		return false
	}
	sum := 1
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			sum += i
			if i*i < num {
				sum += num / i
			}
			if sum > num {
				return false
			}
		}
	}
	return sum == num
}
