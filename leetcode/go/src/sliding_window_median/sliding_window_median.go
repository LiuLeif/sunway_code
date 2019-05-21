// 2017-11-29 17:57
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
	BottomHeap MaxHeap
	TopHeap    MinHeap
}

/** initialize your data structure here. */
func NewMdianFinder() MedianFinder {
	x := MedianFinder{}
	heap.Init(&(x.BottomHeap))
	heap.Init(&(x.TopHeap))
	return x
}

func (this *MedianFinder) Balance() {
	if len(this.TopHeap)-len(this.BottomHeap) > 1 {
		x := heap.Pop(&(this.TopHeap))
		heap.Push(&(this.BottomHeap), x)
	} else if len(this.BottomHeap) > len(this.TopHeap) {
		x := heap.Pop(&(this.BottomHeap))
		heap.Push(&(this.TopHeap), x)
	}
}

func (this *MedianFinder) RemoveNum(num int) {
	if num >= this.TopHeap[0] {
		for i := 0; i < len(this.TopHeap); i++ {
			if this.TopHeap[i] == num {
				heap.Remove(&this.TopHeap, i)
				break
			}
		}
	} else {
		for i := 0; i < len(this.BottomHeap); i++ {
			if this.BottomHeap[i] == num {
				heap.Remove(&this.BottomHeap, i)
				break
			}
		}
	}
	this.Balance()
}

func (this *MedianFinder) AddNum(num int) {
	if len(this.TopHeap) == 0 || num >= this.TopHeap[0] {
		heap.Push(&(this.TopHeap), num)
	} else {
		heap.Push(&(this.BottomHeap), num)
	}
	this.Balance()
}

func (this *MedianFinder) FindMedian() float64 {
	if len(this.BottomHeap) == len(this.TopHeap) {
		return float64(this.BottomHeap[0]+this.TopHeap[0]) / 2
	} else {
		return float64(this.TopHeap[0])
	}
}

func medianSlidingWindow(nums []int, k int) []float64 {
	ret := []float64{}

	finder := NewMdianFinder()
	for i := 0; i < k; i++ {
		finder.AddNum(nums[i])
	}

	for i := k; i < len(nums); i++ {
		ret = append(ret, finder.FindMedian())
		finder.RemoveNum(nums[i-k])
		finder.AddNum(nums[i])
	}
	ret = append(ret, finder.FindMedian())
	return ret
}
func main() {
	fmt.Println(medianSlidingWindow([]int{9, 7, 0, 3, 9, 8, 6, 5, 7, 6}, 2))
}
