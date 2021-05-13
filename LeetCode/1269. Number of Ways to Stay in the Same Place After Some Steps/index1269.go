package main

import "fmt"

func main() {
	fmt.Println(numWays(500, 969997)) // 51
	fmt.Println(numWays(6, 13))       // 51
	fmt.Println(numWays(27, 7))       // 127784505
	fmt.Println(numWays(3, 2))        // 4
	fmt.Println(numWays(2, 4))        // 2
	fmt.Println(numWays(4, 2))        // 8
}

func numWays(steps int, arrLen int) int {
	const mod = 1000000007
	if arrLen > steps/2+1 {
		arrLen = steps/2 + 1
	}
	dpNow := make([]int, arrLen)
	dpNow[0] = 1

	for i := 1; i <= steps; i++ {
		tmp := dpNow[0]
		left := arrLen
		if (steps-i/2)+1 < arrLen {
			left = (steps - i/2) + 1
		}
		for j := 0; j < left; j++ {
			cur := dpNow[j]
			if j-1 >= 0 {
				dpNow[j] = (dpNow[j] + tmp) % mod
			}
			if j+1 < arrLen {
				dpNow[j] = (dpNow[j] + dpNow[j+1]) % mod
			}
			tmp = cur
		}
	}

	return dpNow[0]
}
