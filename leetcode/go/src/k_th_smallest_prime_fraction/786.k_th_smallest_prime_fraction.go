// 2018-12-13 17:42
package main

import (
	"container/heap"
	"fmt"
)

type Pair struct {
	a, b int
}
type MinHeap struct {
	data  []Pair
	array []int
}

func (h MinHeap) Len() int {
	return len(h.data)
}

func (h MinHeap) Less(i, j int) bool {
	p1, p2 := h.data[i], h.data[j]
	return h.array[p1.a]*h.array[p2.b] < h.array[p1.b]*h.array[p2.a]
}

func (h MinHeap) Swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func (ph *MinHeap) Pop() interface{} {
	t := (ph.data)[len(ph.data)-1]
	ph.data = (ph.data)[:len(ph.data)-1]
	return t
}

func (ph *MinHeap) Push(x interface{}) {
	ph.data = append(ph.data, x.(Pair))
}

func kthSmallestPrimeFraction(A []int, K int) []int {
	// 1 2 3 4 5
	N := len(A)
	queue := MinHeap{}
	queue.data = []Pair{}
	queue.array = A

	for i := 0; i < len(A)-1; i++ {
		heap.Push(&queue, Pair{i, N - 1})
	}
	for i := K; i > 1; i-- {
		top := heap.Pop(&queue).(Pair)
		if top.b > top.a+1 {
			heap.Push(&queue, Pair{top.a, top.b - 1})
		}
	}
	top := heap.Pop(&queue).(Pair)
	return []int{A[top.a], A[top.b]}
}

func main() {
	fmt.Println(kthSmallestPrimeFraction([]int{1, 2, 3, 5}, 3))
	fmt.Println(kthSmallestPrimeFraction([]int{1, 7}, 1))
}
