// 2017-12-14 16:21
package main

import (
	"container/heap"
	"fmt"
)

type Pair struct {
	s     string
	count int
}

type MinHeap []Pair

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	if h[i].count < h[j].count {
		return false
	}
	if h[i].count > h[j].count {
		return true
	}
	return h[i].s < h[j].s
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
	*ph = append(*ph, x.(Pair))
}

func topKFrequent(words []string, k int) []string {
	ret := []string{}
	m := make(map[string]int)
	for _, v := range words {
		m[v]++
	}

	pairs := make(MinHeap, len(m))
	i := 0
	for k, v := range m {
		pairs[i] = Pair{k, v}
		i++
	}

	heap.Init(&pairs)
	for i := 0; i < k; i++ {
		ret = append(ret, heap.Pop(&pairs).(Pair).s)
	}

	return ret
}
func main() {
	fmt.Println(topKFrequent([]string{"i", "love", "leetcode", "i", "love", "coding"}, 1))
}
