package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(compareVersion("0.1", "1.1"))       // -1
	fmt.Println(compareVersion("1.0.1", "1"))       // 1
	fmt.Println(compareVersion("7.5.2.4", "7.5.3")) //-1
	fmt.Println(compareVersion(".1", "0.001"))      //0
	fmt.Println(compareVersion("1.01", "1.001"))    //0
	fmt.Println(compareVersion("1.0", "1.0.0"))     //0
}

func compareVersion(version1 string, version2 string) int {
	v1 := strings.Split(version1, ".")
	v2 := strings.Split(version2, ".")

	swaped := 1
	if len(v2) > len(v1) {
		v1, v2 = v2, v1
		swaped = -1
	}

	for i := 0; i < len(v1); i++ {
		i1, _ := strconv.Atoi(v1[i])
		if i >= len(v2) {
			if i1 > 0 {
				return swaped
			}
		} else {
			i2, _ := strconv.Atoi(v2[i])
			if i1 > i2 {
				return swaped
			} else if i1 < i2 {
				return swaped * -1
			}
		}
	}
	return 0
}
