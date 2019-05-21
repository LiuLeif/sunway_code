// 2018-11-07 14:27
package main

import (
	"container/heap"
	"fmt"
	"math"
	"sort"
)

type Pair struct {
	ratio   float64
	quality int
}

type MaxHeap []Pair

func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Less(i, j int) bool {
	return h[i].quality > h[j].quality
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

func mincostToHireWorkers(quality []int, wage []int, K int) float64 {

	N := len(quality)
	factor := make([]Pair, N)
	for i := 0; i < N; i++ {
		factor[i] = Pair{float64(wage[i]) / float64(quality[i]), quality[i]}
	}

	sort.Slice(factor, func(i, j int) bool {
		return factor[i].ratio < factor[j].ratio
	})

	queue := MaxHeap{}
	qualityTotal := 0
	ret := 0.0
	for i := 0; i < K; i++ {
		qualityTotal += factor[i].quality
		heap.Push(&queue, factor[i])
	}
	ret = float64(qualityTotal) * factor[K-1].ratio
	for i := K; i < N; i++ {
		top := heap.Pop(&queue).(Pair)
		qualityTotal -= top.quality
		qualityTotal += factor[i].quality
		heap.Push(&queue, factor[i])
		ret = math.Min(ret, float64(qualityTotal)*factor[i].ratio)
	}
	return ret
}

func main() {
	fmt.Println(mincostToHireWorkers(
		[]int{3, 1, 10, 10, 1},
		[]int{4, 8, 2, 2, 7}, 3))
	fmt.Println(mincostToHireWorkers(
		[]int{3, 4, 3},
		[]int{13, 8, 20}, 1))
}
