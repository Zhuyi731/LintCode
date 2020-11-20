package main

import "fmt"

func main() {
	obj := Constructor()
	obj.Add(1)
	obj.Add(2)
	fmt.Println(obj.Contains(1))
	fmt.Println(obj.Contains(3))
	obj.Add(2)
	fmt.Println(obj.Contains(2))
	obj.Remove(2)
	fmt.Println(obj.Contains(2))
}

type MyHashSet struct {
	storage []int
}

/** Initialize your data structure here. */
func Constructor() MyHashSet {
	storage := make([]int, 1000001)
	hashSet := MyHashSet{
		storage: storage,
	}
	return hashSet
}

func (this *MyHashSet) Add(key int) {
	this.storage[key] = 1
}

func (this *MyHashSet) Remove(key int) {
	this.storage[key] = 0
}

/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
	return this.storage[key] == 1
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
