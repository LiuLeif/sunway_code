// 2018-11-07 13:58
package main

import (
	"container/heap"
	"fmt"
)

type Tuple struct {
	count int
	index int
	value int
}

type MaxHeap []Tuple

func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Less(i, j int) bool {
	if h[i].count > h[j].count {
		return true
	}
	if h[i].count < h[j].count {
		return false
	}

	return h[i].index > h[j].index
}

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (ph *MaxHeap) Pop() interface{} {
	t := (*ph)[len(*ph)-1]
	*ph = (*ph)[:len(*ph)-1]
	return t
}

func (ph *MaxHeap) Push(x interface{}) {
	*ph = append(*ph, x.(Tuple))
}

type FreqStack struct {
	heap    MaxHeap
	index   int
	counter map[int]int
}

func Constructor() FreqStack {
	return FreqStack{MaxHeap{}, 0, map[int]int{}}
}

func (this *FreqStack) Push(x int) {
	this.counter[x]++
	this.index++
	heap.Push(&this.heap, Tuple{this.counter[x], this.index, x})
}

func (this *FreqStack) Pop() int {
	top := heap.Pop(&this.heap).(Tuple)
	this.counter[top.value]--
	return top.value
}

func main() {
	stack := Constructor()
	stack.Push(5)
	stack.Push(7)
	stack.Push(5)
	stack.Push(7)
	stack.Push(4)
	stack.Push(5)

	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
}
