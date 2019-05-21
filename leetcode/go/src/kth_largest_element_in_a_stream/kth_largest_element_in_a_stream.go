// 2018-10-12 18:16
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

type KthLargest struct {
	theHeap MinHeap
	k       int
}

func Constructor(k int, nums []int) KthLargest {
	theHeap := MinHeap(nums)
	heap.Init(&theHeap)
	ret := KthLargest{theHeap, k}
	return ret
}

func (this *KthLargest) Add(val int) int {
	heap.Push(&(this.theHeap), val)
	for len(this.theHeap) > this.k {
		heap.Pop(&(this.theHeap))
	}
	return this.theHeap[0]
}

func main() {
	kthLargest := Constructor(3, []int{1, 2, 3})
	fmt.Println(kthLargest.Add(3))
	fmt.Println(kthLargest.Add(3))
	fmt.Println(kthLargest.Add(4))
	fmt.Println(kthLargest.Add(5))
	fmt.Println(kthLargest.Add(6))
}
