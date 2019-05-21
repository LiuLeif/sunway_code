// 2018-01-10 11:13
package main

import (
	"container/heap"
	"fmt"
	"math"
)

type Neightbor struct {
	i, w int
}

func max(nums ...int) int {
	ret := math.MinInt32
	for _, n := range nums {
		if ret < n {
			ret = n
		}
	}
	return ret
}

func min(nums ...int) int {
	ret := math.MaxInt32
	for _, n := range nums {
		if ret > n {
			ret = n
		}
	}
	return ret
}

type MinHeap []Pair

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i].d < h[j].d
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

type Pair struct {
	d, i int
}

func networkDelayTime(times [][]int, N int, K int) int {
	neighbors := map[int][]Neightbor{}
	for i := 0; i < len(times); i++ {
		f, t, w := times[i][0], times[i][1], times[i][2]
		if neighbors[f] == nil {
			neighbors[f] = []Neightbor{}
		}
		neighbors[f] = append(neighbors[f], Neightbor{t, w})
	}
	queue := MinHeap([]Pair{{0, K}})
	heap.Init(&queue)

	distance := map[int]int{}
	for i := 1; i < N+1; i++ {
		distance[i] = math.MaxInt32
	}
	distance[K] = 0

	visited := map[int]bool{}

	for len(queue) != 0 {
		top := heap.Pop(&queue).(Pair).i
		if neighbors[top] == nil {
			continue
		}
		if visited[top] {
			continue
		}
		for i := 0; i < len(neighbors[top]); i++ {
			neighbor := neighbors[top][i]
			if distance[neighbor.i] > distance[top]+neighbor.w {
				distance[neighbor.i] = distance[top] + neighbor.w
				heap.Push(&queue, Pair{distance[neighbor.i], neighbor.i})
			}
		}
		visited[top] = true
	}
	ret := 0
	for i := 1; i < N+1; i++ {
		ret = max(ret, distance[i])
	}
	if ret == math.MaxInt32 {
		return -1
	}
	return ret
}
func main() {
	//             1
	//            / \
	//           / . \
	//        4 /     \ 2
	//         v        v
	//        2<---------3
	//              1
	fmt.Println(networkDelayTime(
		[][]int{
			{1, 2, 4},
			{1, 3, 2},
			{3, 2, 1},
		}, 3, 1,
	))
}
