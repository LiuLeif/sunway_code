// 2018-11-01 16:59
package main

import (
	"container/heap"
	"fmt"
)

type MinHeap []int

func (this MinHeap) Len() int {
	return len(this)
}

func (this MinHeap) Less(i, j int) bool {
	return this[i] < this[j]
}

func (this *MinHeap) Swap(i, j int) {
	(*this)[i], (*this)[j] = (*this)[j], (*this)[i]
}

func (this *MinHeap) Push(x interface{}) {
	*this = append(*this, x.(int))
}

func (this *MinHeap) Pop() interface{} {
	ret := (*this)[len(*this)-1]
	(*this) = (*this)[:len(*this)-1]
	return ret
}

func nthSuperUglyNumber(n int, primes []int) int {
	queue := MinHeap{}
	heap.Init(&queue)
	heap.Push(&queue, 1)
	prev := 0
	for {
		top := heap.Pop(&queue).(int)
		if prev == top {
			continue
		}
		prev = top
		n--
		if n == 0 {
			return top
		}
		for _, v := range primes {
			heap.Push(&queue, v*top)
		}
	}
	return 0
}

func main() {
	// 2 7 13 19
	// 1 2 4 7 8 13 14 16
	fmt.Println(nthSuperUglyNumber(12, []int{2, 7, 13, 19}))
}
