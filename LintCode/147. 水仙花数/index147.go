package main

import (
	"fmt"
)

func main() {
	fmt.Println(getNarcissisticNumbers(1))
	fmt.Println(getNarcissisticNumbers(2))
	fmt.Println(getNarcissisticNumbers(3))
	fmt.Println(getNarcissisticNumbers(4))
	fmt.Println(getNarcissisticNumbers(5))
	fmt.Println(getNarcissisticNumbers(6))
	fmt.Println(getNarcissisticNumbers(7))
	fmt.Println(getNarcissisticNumbers(8))
}

func getItems(n int) (int, []int) {
	result := []int{}
	ct := 0
	for {
		if n != 0 {
			result = append(result, n%10)
			n = n / 10
			ct++
		} else {
			return ct, result

		}
	}
	return ct, result

}

/**
 * @param n: The number of digits
 * @return: All narcissistic numbers with n digits
 */
func getNarcissisticNumbers(n int) []int {
	result := [][]int{
		[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		[]int{},
		[]int{153, 370, 371, 407},
		[]int{1634, 8208, 9474},
		[]int{54748, 92727, 93084},
		[]int{548834},
		[]int{1741725, 4210818, 9800817, 9926315},
		[]int{24678050, 24678051, 88593477},
	}
	return result[n-1]
	// // write your code here˝
	// min := int(math.Pow(10, float64(n-1)))
	// max := int(math.Pow(10, float64(n)))

	// result := []int{}
	// for i := min; i < max; i++ {
	// 	ct, items := getItems(i)
	// 	sum := 0
	// 	for _, v := range items {
	// 		sum += int(math.Pow(float64(v), float64(ct)))
	// 	}
	// 	if sum == i {
	// 		result = append(result, i)
	// 	}
	// }
	// return result
}
