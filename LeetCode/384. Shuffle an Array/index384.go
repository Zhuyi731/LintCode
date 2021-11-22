package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	s := Constructor([]int{3, 1, 2})
	fmt.Println(s.Shuffle())
	fmt.Println(s.Reset())
	fmt.Println(s.Shuffle())
}

type Solution struct {
	Origin  []int
	Current []int
	Len     int
}

func Constructor(nums []int) Solution {
	return Solution{
		Origin:  nums,
		Current: append([]int{}, nums...),
		Len:     len(nums),
	}
}

func (this *Solution) Reset() []int {
	for i := range this.Origin {
		this.Current[i] = this.Origin[i]
	}
	return this.Origin
}

func (this *Solution) Shuffle() []int {
	for i, _ := range this.Current {
		j := randomNum(0, this.Len)
		this.Current[i], this.Current[j] = this.Current[j], this.Current[i]
	}
	return this.Current
}
func init() {
	rand.Seed(time.Now().Unix())
}
func randomNum(start, end int) int {
	l := end - start
	r := rand.Intn(l)
	return r + start
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
