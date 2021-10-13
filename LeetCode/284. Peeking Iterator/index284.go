/*   Below is the interface for Iterator, which is already defined for you.
 *
 *   type Iterator struct {
 *
 *   }
 *
 *   func (this *Iterator) hasNext() bool {
 *		// Returns true if the iteration has more elements.
 *   }
 *
 *   func (this *Iterator) next() int {
 *		// Returns the next element in the iteration.
 *   }
 */
package main

import "fmt"

func main() {
	actions([]string{"PeekingIterator", "next", "peek", "next", "next", "hasNext"}, [][]int{{1, 2, 3}})
}

func actions(act []string, params [][]int) {
	var p *PeekingIterator
	for i, v := range act {
		switch v {
		case "PeekingIterator":
			i := Iterator{
				Data:         params[i],
				Len:          len(params[i]),
				CurrentIndex: 0,
			}
			p = Constructor(&i)
		case "next":
			fmt.Println(p.next())
		case "peek":
			fmt.Println(p.peek())
		}
	}
}

type Iterator struct {
	Data         []int
	CurrentIndex int
	Len          int
}

func (this *Iterator) hasNext() bool {
	// Returns true if the iteration has more elements.
	return this.CurrentIndex < this.Len
}

func (this *Iterator) next() int {
	// Returns the next element in the iteration.
	val := this.Data[this.CurrentIndex]
	this.CurrentIndex++
	return val
}

type PeekingIterator struct {
	Iter    *Iterator
	HasNext bool
	NextVal int
}

func Constructor(iter *Iterator) *PeekingIterator {
	next := -1
	if iter.hasNext() {
		next = iter.next()
	}
	return &PeekingIterator{
		Iter:    iter,
		HasNext: iter.hasNext(),
		NextVal: next,
	}
}

func (this *PeekingIterator) hasNext() bool {
	return this.HasNext
}

func (this *PeekingIterator) next() int {
	next := this.NextVal
	this.HasNext = this.Iter.hasNext()
	if this.HasNext {
		this.NextVal = this.Iter.next()
	}
	return next
}

func (this *PeekingIterator) peek() int {
	return this.NextVal
}
