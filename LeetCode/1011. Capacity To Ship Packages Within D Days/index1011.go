package main

import "fmt"

func main() {
	fmt.Println(shipWithinDays([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 1))
	fmt.Println(shipWithinDays([]int{1, 2, 3, 1, 1}, 4))
	fmt.Println(shipWithinDays([]int{3, 2, 2, 4, 1, 4}, 3))
	fmt.Println(shipWithinDays([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5))
}

func shipWithinDays(weights []int, D int) int {
	sum := 0
	avg := 0
	max := 0
	for i := 0; i < len(weights); i++ {
		sum += weights[i]
		if weights[i] > max {
			max = weights[i]
		}
	}
	avg = sum / D // 至少为这个
	if max > avg {
		avg = max
	}

	l := avg
	r := sum
	mid := (l + r) / 2

	isMidOk := func(mid int) bool {
		ct := 1
		curSum := 0
		for _, v := range weights {
			if curSum+v <= mid {
				curSum = curSum + v
			} else {
				ct++
				curSum = v
				if ct > D {
					return false
				}
			}
		}
		return true
	}

	// 假设为mid
	for {
		if l+1 >= r {
			if isMidOk(l) {
				return l
			}
			return r
		}
		ok := isMidOk(mid)
		if ok { // 往左
			r = mid
			mid = (l + r) / 2
		} else {
			l = mid
			mid = (l + r) / 2
		}
	}
}
