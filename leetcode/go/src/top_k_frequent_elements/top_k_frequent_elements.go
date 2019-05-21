// 2017-11-17 14:21
package main

import (
	"container/heap"
	"fmt"
)

type Pair struct {
	First, Second int
}

type MaxHeap []Pair

func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Less(i, j int) bool {
	return h[i].First > h[j].First
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
	*ph = append(*ph, x.(Pair))
}

func topKFrequent(nums []int, k int) []int {
	m := map[int]int{}
	for _, v := range nums {
		m[v]++
	}

	h := make(MaxHeap, len(m))
	i := 0
	for k, v := range m {
		h[i] = Pair{v, k}
		i++
	}

	heap.Init(&h)

	ret := make([]int, k)
	for i := 0; i < k; i++ {
		p := heap.Pop(&h)
		ret[i] = p.(Pair).Second
	}

	return ret
}
func main() {
	fmt.Println(topKFrequent([]int{1, 2, 2, 3}, 2))
}
