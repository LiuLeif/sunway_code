// 2017-12-29 16:35
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

func nthUglyNumber(n int) int {
	x := MinHeap{1}
	heap.Init(&x)
	count := 0
	m := map[int]bool{}
	m[1] = true
	top := 0
	for count < n {
		count++
		top = heap.Pop(&x).(int)
		for _, v := range []int{2, 3, 5} {
			if !m[top*v] {
				heap.Push(&x, top*v)
				m[top*v] = true
			}
		}
	}

	return top
}
func main() {
	// 1, 2, 3, 4, 5, 6, 8, 9, 10, 12
	fmt.Println(nthUglyNumber(10))
}
