package main

import (
	"fmt"
	"strconv"
)

func main() {

	fmt.Println(simplifiedFractions(1))
	fmt.Println(simplifiedFractions(2))
	fmt.Println(simplifiedFractions(3))
	fmt.Println(simplifiedFractions(4))
	fmt.Println(simplifiedFractions(5))
	fmt.Println(simplifiedFractions(6))
}

func simplifiedFractions(n int) []string {
	result := []string{}
	for i := 2; i <= n; i++ {
		result = append(result, getN(i)...)
	}
	return result
}

func getN(n int) []string {
	if n == 1 {
		return []string{}
	}

	facts := []int{}
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			facts = append(facts, i)
		}
	}
	subfix := "/" + strconv.Itoa(n)
	result := []string{"1" + subfix}
	for i := 2; i < n; i++ {
		skip := false
		for j := 0; j < len(facts); j++ {
			if i%facts[j] == 0 {
				skip = true
			}
		}
		if !skip {
			result = append(result, strconv.Itoa(i)+subfix)
		}
	}
	return result
}
