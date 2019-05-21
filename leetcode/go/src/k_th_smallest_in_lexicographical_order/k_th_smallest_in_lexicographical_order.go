// 2018-01-03 14:30
package main

import (
	"container/heap"
	"fmt"
	"math"
	"strconv"
)

func max(nums ...int) int {
	ret := math.MinInt32
	for _, n := range nums {
		if ret < n {
			ret = n
		}
	}
	return ret
}

type MinHeap []string

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	x, y := h[i], h[j]
	lx, ly := len(x), len(y)
	len := max(lx, ly)
	for i := 0; i < len; i++ {
		a, b := byte(0), byte(0)
		if i < lx {
			a = x[i]
		}
		if i < ly {
			b = y[i]
		}
		if a < b {
			return true
		}
		if a > b {
			return false
		}
	}
	return false
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
	*ph = append(*ph, x.(string))
}

func findKthNumber(n int, k int) int {
	nums := make([]string, n)
	for i := 0; i < n; i++ {
		nums[i] = strconv.Itoa(i + 1)
	}
	h := MinHeap(nums)
	heap.Init(&h)

	ret := 0
	for i := 0; i < k; i++ {
		ret, _ = strconv.Atoi(heap.Pop(&h).(string))
	}

	return ret
}
func main() {
	// n=13
	// 1,10,11,12,2,3,4,5,6,7,8,9
	// LTE
	fmt.Println(findKthNumber(1300, 2))
}
