package main

/**
开启3个go routine
循环打印dog fish cat 100次
*/

import "fmt"

type chanList struct {
	name string
	c    chan int
}

var quitCh chan int
var ct int

func main() {
	quitCh = make(chan int)
	ct = 0
	cList := []chanList{
		{name: "dog", c: make(chan int)},
		{name: "fish", c: make(chan int)},
		{name: "cat", c: make(chan int)},
		{name: "cat2", c: make(chan int)},
		{name: "cat23", c: make(chan int)},
		{name: "cat234", c: make(chan int)},
		{name: "cat2345", c: make(chan int)},
	}

	for _, v := range cList {
		go run(v.name, cList)
	}
	<-quitCh
}

func findSelfChan(name string, cList []chanList) int {
	for i, v := range cList {
		if v.name == name {
			return i
		}
	}
	return -1
}

func run(name string, cList []chanList) {
	for {
		l := len(cList)
		i := findSelfChan(name, cList)
		if i == 0 {
			if ct > 0 {
				prev := (i - 1 + l) % l
				<-cList[prev].c
			}
			ct++
		} else {
			prev := (i - 1 + l) % l
			<-cList[prev].c
		}
		fmt.Println(name)
		cList[i].c <- 1
		if i == len(cList)-1 && ct == 100 {
			quitCh <- 1
			break
		}
	}
}
