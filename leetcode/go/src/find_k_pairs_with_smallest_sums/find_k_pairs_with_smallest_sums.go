// 2017-12-08 14:35
package main

import (
	"container/heap"
	"fmt"
	"math"
)

type MinHeap []Tuple

func (heap MinHeap) Len() int {
	return len(heap)
}

func (heap MinHeap) Swap(i, j int) {
	heap[i], heap[j] = heap[j], heap[i]
}

func (heap MinHeap) Less(i, j int) bool {
	return heap[i].val < heap[j].val
}

func (heap *MinHeap) Push(a interface{}) {
	*heap = append(*heap, a.(Tuple))
}

func (heap *MinHeap) Pop() interface{} {
	ret := (*heap)[len(*heap)-1]
	*heap = (*heap)[:len(*heap)-1]
	return ret
}

type Tuple struct {
	i, j, val int
}

//   2 4 6
// 1
// 7
// 11

func min(nums ...int) int {
	ret := math.MaxInt32
	for _, n := range nums {
		if ret > n {
			ret = n
		}
	}
	return ret
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	ret := [][]int{}
	if len(nums1) == 0 || len(nums2) == 0 || k <= 0 {
		return ret
	}
	minHeap := MinHeap{}
	for j := 0; j < len(nums2); j++ {
		heap.Push(&minHeap, Tuple{0, j, nums1[0] + nums2[j]})
	}

	for i := 0; i < min(k, len(nums1)*len(nums2)); i++ {
		top := heap.Pop(&minHeap).(Tuple)
		ret = append(ret, []int{nums1[top.i], nums2[top.j]})
		if top.i < len(nums1)-1 {
			heap.Push(&minHeap, Tuple{top.i + 1, top.j, nums1[top.i+1] + nums2[top.j]})
		}
	}

	return ret
}
func main() {
	fmt.Println(kSmallestPairs([]int{1, 2, 3}, []int{1, 2, 3}, 10))
}
