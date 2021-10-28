package main

import "fmt"

func main() {
	fmt.Println(reorderedPowerOf2(1))
	fmt.Println(reorderedPowerOf2(10))
	fmt.Println(reorderedPowerOf2(16))
	fmt.Println(reorderedPowerOf2(24))
	fmt.Println(reorderedPowerOf2(46))
}

func init() {
	c := []int{}
	for i := 0; i <= 32; i++ {
		c = append(c, 1<<i)
	}
	for _, v := range c {
		mp := map[int]int{}
		for v != 0 {
			mp[v%10]++
			v = v / 10
		}
		total = append(total, mp)
	}
	fmt.Println(total)
}

var total []map[int]int

func reorderedPowerOf2(n int) bool {
	mp := map[int]int{}
	for n != 0 {
		mp[n%10]++
		n = n / 10
	}
	for _, v := range total {
		flag := 0
		for key, value := range v {
			if mp[key] != value {
				flag = 1
				break
			}
		}
		for key, value := range mp {
			if v[key] != value {
				flag = 1
				break
			}
		}
		if flag == 0 {
			return true
		}
	}
	return false
}
