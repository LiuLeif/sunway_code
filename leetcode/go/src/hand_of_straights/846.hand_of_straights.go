// 2018-10-25 17:29
package main

import "fmt"
import "container/heap"
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

func isNStraightHand(hand []int, W int) bool {
	counter := map[int]int{}
	for _, v := range hand {
		counter[v]++
	}
	keys := []int{}
	for k, _ := range counter {
		keys = append(keys, k)
	}
	minHeap := MinHeap(keys)
	heap.Init(&minHeap)
	for len(minHeap) != 0 {
		top := heap.Pop(&minHeap).(int)
		count := counter[top]
		if count == 0 {
			continue
		}
		for i := 0; i < W; i++ {
			if counter[top+i] < count {
				return false
			}
			counter[top+i] -= count
		}
	}
	return true
}

func main() {
	fmt.Println(isNStraightHand([]int{1, 2, 3, 4, 5}, 5))
}
