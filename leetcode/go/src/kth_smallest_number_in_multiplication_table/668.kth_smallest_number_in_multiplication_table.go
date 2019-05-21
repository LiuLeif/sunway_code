// 2018-12-12 11:09
package main

import (
	"fmt"
	"math"
)

// type Tuple struct {
// 	v, row, col int
// }

// type MinHeap []Tuple

// func (h MinHeap) Len() int {
// 	return len(h)
// }

// func (h MinHeap) Less(i, j int) bool {
// 	return h[i].v < h[j].v
// }

// func (h MinHeap) Swap(i, j int) {
// 	h[i], h[j] = h[j], h[i]
// }

// func (ph *MinHeap) Pop() interface{} {
// 	t := (*ph)[len(*ph)-1]
// 	*ph = (*ph)[:len(*ph)-1]
// 	return t
// }

// func (ph *MinHeap) Push(x interface{}) {
// 	*ph = append(*ph, x.(Tuple))
// }

// func findKthNumber(m int, n int, k int) int {
// 	queue := MinHeap{}
// 	for i := 0; i < n; i++ {
// 		heap.Push(&queue, Tuple{i + 1, 0, i})
// 	}
// 	for i := 0; i < k-1; i++ {
// 		top := heap.Pop(&queue).(Tuple)
// 		if top.row < m-1 {
// 			heap.Push(&queue, Tuple{top.v + top.col + 1, top.row + 1, top.col})
// 		}
// 	}
// 	return heap.Pop(&queue).(Tuple).v
// }

func findKthNumber(m int, n int, k int) int {
	min := func(nums ...int) int {
		ret := math.MaxInt32
		for _, n := range nums {
			if ret > n {
				ret = n
			}
		}
		return ret
	}

	count := func(x int) int {
		ret := 0
		for i := 0; i < m; i++ {
			ret += min(x/(i+1), n)
		}
		return ret
	}

	lo, hi := 1, m*n
	for lo <= hi {
		mid := (lo + hi) / 2
		if count(mid) < k {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return lo
}

func main() {
	fmt.Println(findKthNumber(2, 3, 6))
	fmt.Println(findKthNumber(27593, 21500, 498591911))
}
