// 2017-11-14 18:54
package main

import (
	"container/heap"
	"fmt"
)

type MinHeap []int

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (ph *MinHeap) Pop() interface{} {
	t := (*ph)[len(*ph)-1]
	*ph = (*ph)[:len(*ph)-1]
	return t
}

func (ph *MinHeap) Push(x interface{}) {
	*ph = append(*ph, x.(int))
}

type MaxHeap []int

func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Less(i, j int) bool {
	return h[i] > h[j]
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
	*ph = append(*ph, x.(int))
}

type MedianFinder struct {
	UpHeap     MaxHeap
	BottomHeap MinHeap
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	x := MedianFinder{}
	heap.Init(&(x.UpHeap))
	heap.Init(&(x.BottomHeap))
	return x
}

func (this *MedianFinder) AddNum(num int) {
	if len(this.BottomHeap) == 0 || num >= this.BottomHeap[0] {
		heap.Push(&(this.BottomHeap), num)
		if len(this.BottomHeap)-len(this.UpHeap) > 1 {
			x := heap.Pop(&(this.BottomHeap))
			heap.Push(&(this.UpHeap), x)
		}
	} else {
		heap.Push(&(this.UpHeap), num)
		if len(this.UpHeap) > len(this.BottomHeap) {
			x := heap.Pop(&(this.UpHeap))
			heap.Push(&(this.BottomHeap), x)
		}
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if len(this.UpHeap) == len(this.BottomHeap) {
		return float64(this.UpHeap[0]+this.BottomHeap[0]) / 2
	} else {
		return float64(this.BottomHeap[0])
	}
}

func main() {
	x := Constructor()
	x.AddNum(1)
	x.AddNum(2)
	x.AddNum(3)
	fmt.Println(x.FindMedian())
}
