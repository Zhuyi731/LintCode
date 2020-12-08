package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(splitIntoFibonacci("3611537383985343591834441270352104793375145479938855071433500231900737525076071514982402115895535257195564161509167334647108949738176284385285234123461518508746752631120827113919550237703163294909")) // []
	fmt.Println(splitIntoFibonacci("539834657215398346785398346991079669377161950407626991734534318677529701785098211336528511"))                                                                                                           // []
	fmt.Println(splitIntoFibonacci("1320581321313221264343965566089105744171833277577"))                                                                                                                                                    // []
	fmt.Println(splitIntoFibonacci("539834657215398346785398346991079669377161950407626991734534318677529701785098211336528511"))                                                                                                           // []
	fmt.Println(splitIntoFibonacci("0123"))                                                                                                                                                                                                 // []
	fmt.Println(splitIntoFibonacci("000"))                                                                                                                                                                                                  // []
	fmt.Println(splitIntoFibonacci("0000"))                                                                                                                                                                                                 // []
	fmt.Println(splitIntoFibonacci("1230456579"))                                                                                                                                                                                           // []
	fmt.Println(splitIntoFibonacci("112358130"))                                                                                                                                                                                            // []]
	fmt.Println(splitIntoFibonacci("123456579"))                                                                                                                                                                                            // [123,456,579]
	fmt.Println(splitIntoFibonacci("11235813"))                                                                                                                                                                                             // [1,1,2,3,5,8,13]
	fmt.Println(splitIntoFibonacci("1101111"))                                                                                                                                                                                              // [110, 1, 111]

}

func splitIntoFibonacci(S string) (result []int) {
	lens := len(S)
	max := 0x7fffffff - 1
	for i := 1; i <= lens/2; i++ {
		// 从一位长度算起，往后推
		if i != 1 && S[0] == '0' {
			break
		}
		first, _ := strconv.Atoi(S[:i])
		for j := 1; j <= lens/2; j++ {
			if j+i >= lens || (j != 1 && S[i] == '0') {
				break
			}
			second, _ := strconv.Atoi(S[i : i+j])
			result := []int{first, second}
			Sleft := S[i+j:]
			if first > max || second > max {
				break
			}
			if judge(Sleft, first, second, &result) {
				return result
			}
		}
	}
	return
}

func judge(S string, first, second int, result *[]int) bool {
	val := first + second
	if val > 0x7fffffff-1 {
		return false
	}
	valCp := val
	var length int
	if valCp == 0 {
		length = 1
	} else {
		length = 0
	}
	for valCp != 0 {
		length++
		valCp /= 10
	}
	if length > len(S) {
		return false
	}
	if length != 1 && S[0] == '0' {
		return false
	}
	strVal := S[:length]
	strToIntVal, _ := strconv.Atoi(strVal)

	if val == strToIntVal {
		*result = append(*result, val)
		if length == len(S) {
			return true
		}
		return judge(S[length:], second, val, result)
	} else {
		return false
	}
}
