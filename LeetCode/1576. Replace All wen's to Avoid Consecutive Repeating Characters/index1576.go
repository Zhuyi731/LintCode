package main

import "fmt"

func main() {
	fmt.Println(modifyString("??yw?ipkj?"))
	fmt.Println(modifyString("?zs"))
	fmt.Println(modifyString("ubv?w"))
	fmt.Println(modifyString("j?qg??b"))
}

func modifyString(s string) string {
	charset := []byte{'a', 'b', 'c'}
	result := []byte(s)
	for i, v := range result {
		if v == '?' {
			for _, val := range charset {
				if (i-1 >= 0 && val == result[i-1]) || (i+1 < len(s) && val == result[i+1]) {
					continue
				}
				result[i] = val
				break
			}
		} else {
			result[i] = v
		}
	}
	return string(result)
}
