// 2018-11-23 18:31
package main

import (
	"container/heap"
	"fmt"
	"sort"
)

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

type Pair struct {
	first, second int
}

func findMaximizedCapital(k int, W int, Profits []int, Capital []int) int {
	N := len(Capital)
	v := make([]Pair, N)
	for i := 0; i < N; i++ {
		v[i] = Pair{Profits[i], Capital[i]}
	}
	sort.Slice(v, func(i, j int) bool {
		return v[i].second < v[j].second
	})

	index := 0
	queue := MaxHeap{}
	for i := 0; i < k; i++ {
		for index < N && v[index].second <= W {
			heap.Push(&queue, v[index].first)
			index++
		}
		if len(queue) == 0 {
			break
		}
		W += heap.Pop(&queue).(int)
	}
	return W
}

func main() {
	fmt.Println(findMaximizedCapital(1, 0, []int{1, 2, 2}, []int{1, 1, 2}))
}
