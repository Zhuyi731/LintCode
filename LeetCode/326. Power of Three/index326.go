package main

import "fmt"

func main() {
	fmt.Println(isPowerOfThree(27))
	fmt.Println(isPowerOfThree(64))
	fmt.Println(isPowerOfThree(9))
	fmt.Println(isPowerOfThree(0))
	fmt.Println(isPowerOfThree(54))
}

var result = map[int]bool{
	1:          true,
	3:          true,
	9:          true,
	27:         true,
	81:         true,
	243:        true,
	729:        true,
	2187:       true,
	6561:       true,
	19683:      true,
	59049:      true,
	177147:     true,
	531441:     true,
	1594323:    true,
	4782969:    true,
	14348907:   true,
	43046721:   true,
	129140163:  true,
	387420489:  true,
	1162261467: true,
	3486784401: true,
}

func isPowerOfThree(n int) bool {
	return result[n]
}
