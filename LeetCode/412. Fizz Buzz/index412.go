package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(fizzBuzz(15))
}

func fizzBuzz(n int) []string {
	result := make([]string, n)
	for i := 3; i <= n; i += 3 {
		result[i-1] = "Fizz"
	}
	for i := 5; i <= n; i += 5 {
		result[i-1] += "Buzz"
	}
	for i := 0; i < n; i++ {
		if result[i] == "" {
			result[i] = strconv.Itoa(i + 1)
		}
	}
	return result
}
