package main

import (
	"fmt"
	"strconv"
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
	var r1, r2, l1, l2, n1, n2 int
	lenv1 := len(version1)
	lenv2 := len(version2)
	for l1 < lenv1 || l2 < lenv2 {
		n1 = 0
		n2 = 0
		for r1 < lenv1 && version1[r1] != '.' {
			r1++
		}
		if l1 < lenv1 {
			n1, _ = strconv.Atoi(version1[l1:r1])
			l1 = r1 + 1
			r1++
		}

		for r2 < lenv2 && version2[r2] != '.' {
			r2++
		}
		if l2 < lenv2 {
			n2, _ = strconv.Atoi(version2[l2:r2])
			l2 = r2 + 1
			r2++
		}

		if n1 != n2 {
			if n1 > n2 {
				return 1
			}
			return -1
		}
	}

	return 0
}
