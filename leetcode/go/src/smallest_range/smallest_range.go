// 2018-04-13 13:06
package main

import (
	"container/heap"
	"fmt"
	"math"
)

type Pair struct {
	value, index int
}
type MinHeap []Pair

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i].value < h[j].value
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

func smallestRange(nums [][]int) []int {
	minHeap := MinHeap{}
	maxValue := 0
	ret := [2]int{math.MinInt32, math.MaxInt32}
	for i := 0; i < len(nums); i++ {
		heap.Push(&minHeap, Pair{nums[i][0], i})
		if nums[i][0] >= maxValue {
			maxValue = nums[i][0]
		}
		nums[i] = nums[i][1:]
	}
	for {
		minValuePair := heap.Pop(&minHeap).(Pair)
		index := minValuePair.index
		if maxValue-minValuePair.value < ret[1]-ret[0] {
			ret[1] = maxValue
			ret[0] = minValuePair.value
		}
		if len(nums[index]) == 0 {
			break
		}
		nextValue := nums[index][0]
		if nextValue >= maxValue {
			maxValue = nextValue
		}
		heap.Push(&minHeap, Pair{nextValue, index})
		nums[index] = nums[index][1:]
	}
	return ret[:]
}
func main() {
	// 20,24
	fmt.Println(smallestRange([][]int{
		{4, 10, 15, 24, 26},
		{0, 9, 12, 20},
		{5, 18, 22, 30},
	}))
}
