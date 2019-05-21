// 2017-11-24 14:50
package main

import (
	"container/heap"
	"fmt"
	"math"
	"sort"
)

type Heap []int

func (h Heap) Len() int {
	return len(h)
}

func (h Heap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (ph *Heap) Pop() interface{} {
	t := (*ph)[len(*ph)-1]
	*ph = (*ph)[:len(*ph)-1]
	return t
}

func (ph *Heap) Push(x interface{}) {
	*ph = append(*ph, x.(int))
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

func leastInterval(tasks []byte, n int) int {
	counter := make(map[byte]int)
	for i := 0; i < len(tasks); i++ {
		counter[tasks[i]]++
	}

	counts := []int{}
	for _, v := range counter {
		counts = append(counts, v)
	}

	sort.Ints(counts)
	maxCount := counts[len(counts)-1]

	ret := (maxCount - 1) * (n + 1)
	for _, v := range counts {
		if v == maxCount {
			ret++
		}
	}

	return max(ret, len(tasks))
}

func leastIntervalHeap(tasks []byte, n int) int {
	counter := make(map[byte]int)
	for i := 0; i < len(tasks); i++ {
		counter[tasks[i]]++
	}

	counts := []int{}
	for _, v := range counter {
		counts = append(counts, v)
	}

	h := Heap([]int{})

	idle := 0
	for len(counts) != 0 {
		for i := 0; i < len(counts); i++ {
			heap.Push(&h, counts[i])
		}
		counts = []int{}
		k := 0
		for k < n+1 && len(h) != 0 {
			t := heap.Pop(&h).(int)
			t -= 1
			if t != 0 {
				counts = append(counts, t)
			}
			k++
		}
		if len(counts) != 0 {
			idle += n + 1 - k
		}
	}
	return len(tasks) + idle
}

func main() {
	fmt.Println(leastIntervalHeap([]byte{'A', 'A', 'A', 'B', 'B', 'B'}, 50))
}
