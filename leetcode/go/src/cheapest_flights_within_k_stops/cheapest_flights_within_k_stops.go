// 2018-02-27 11:19
package main

import (
	"container/heap"
	"fmt"
	"math"
)

type Pair struct {
	dest int
	cost int
}

type City struct {
	index     int
	heapIndex int
	distance  int
	hop       int
}

type MinHeap []*City

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i].distance < h[j].distance
}

func (h MinHeap) Swap(i, j int) {
	h[i].heapIndex = j
	h[j].heapIndex = i
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Pop() interface{} {
	t := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return t
}

func (h *MinHeap) Push(x interface{}) {
	tmp := x.(*City)
	tmp.heapIndex = len(*h)
	*h = append(*h, tmp)
}

// this code is wrong!
func findCheapestPriceDijkstra(n int, flights [][]int, src int, dst int, K int) int {
	neighbors := map[int][]Pair{}
	for _, flight := range flights {
		src, dst, cost := flight[0], flight[1], flight[2]
		if neighbors[src] == nil {
			neighbors[src] = []Pair{}
		}
		neighbors[src] = append(neighbors[src], Pair{dst, cost})
	}

	cities := map[int]*City{}
	cities[src] = &City{src, 0, 0, 0}
	queue := MinHeap{}
	heap.Push(&queue, cities[src])

	for len(queue) != 0 {
		top := heap.Pop(&queue).(*City)
		for _, p := range neighbors[top.index] {
			if _, ok := cities[p.dest]; !ok {
				cities[p.dest] = &City{p.dest, 0, top.distance + p.cost, top.hop + 1}
				if cities[p.dest].hop <= K {
					heap.Push(&queue, cities[p.dest])
				}
			}
			neighbor := cities[p.dest]
			if neighbor.distance > top.distance+p.cost && neighbor.hop <= K {
				neighbor.distance = top.distance + p.cost
				neighbor.hop = top.hop + 1
				heap.Fix(&queue, neighbor.heapIndex)
			}
		}
	}
	if _, ok := cities[dst]; !ok {
		return -1
	}
	return cities[dst].distance
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

func findCheapestPriceDP(n int, flights [][]int, src int, dst int, K int) int {
	K++
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, K+1)
		for j := 0; j < K+1; j++ {
			dp[i][j] = math.MaxInt32
		}
	}
	dp[src][0] = 0
	for i := 1; i < K+1; i++ {
		for j := 0; j < n; j++ {
			dp[j][i] = dp[j][i-1]
		}
		for _, flight := range flights {
			f, t, cost := flight[0], flight[1], flight[2]
			dp[t][i] = min(dp[t][i], dp[f][i-1]+cost)
		}
	}
	if dp[dst][K] == math.MaxInt32 {
		return -1
	}
	return dp[dst][K]
}

func findCheapestPriceBellman(n int, flights [][]int, src int, dst int, K int) int {
	dis := make([]int, n)
	for i := 0; i < n; i++ {
		dis[i] = math.MaxInt32
	}
	dis[src] = 0
	nextDis := append([]int{}, dis...)
	for i := 0; i <= K; i++ {
		for _, flight := range flights {
			f, t, cost := flight[0], flight[1], flight[2]
			nextDis[t] = min(nextDis[t], dis[f]+cost)
		}
		dis = nextDis
	}
	if dis[dst] == math.MaxInt32 {
		return -1
	}
	return dis[dst]
}

func main() {
	// dijkstra is wrong for this test case:
	fmt.Println(findCheapestPriceBellman(5,
		[][]int{
			{0, 1, 5},
			{1, 2, 5},
			{0, 3, 2},
			{3, 1, 2},
			{1, 4, 1},
			{4, 2, 1},
		}, 0, 2, 2))
}
