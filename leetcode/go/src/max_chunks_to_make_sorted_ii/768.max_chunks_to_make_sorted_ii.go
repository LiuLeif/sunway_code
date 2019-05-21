// 2018-11-14 10:56
package main

import (
	"fmt"
	"math"
)

func maxChunksToSorted(arr []int) int {
	// 2 1 3 4 4
	max := func(nums ...int) int {
		ret := math.MinInt32
		for _, n := range nums {
			if ret < n {
				ret = n
			}
		}
		return ret
	}

	N := len(arr)
	lastSmaller := make([]int, N)
	currMin := math.MaxInt32
	for i := N - 1; i >= 0; i-- {
		if arr[i] < currMin {
			currMin = arr[i]
		}
		lastSmaller[i] = currMin
	}
	// 2 1 3 4 4
	// 1 1 3 4 4
	firstLarger := func(x int) int {
		lo, hi := 0, N-1
		for lo <= hi {
			mid := (lo + hi) / 2
			if lastSmaller[mid] < x {
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		}
		return lo
	}
	ret := 0
	right := 0
	for i := 0; i < N; i++ {
		n := arr[i]
		lastSmallerIndex := firstLarger(n) - 1
		right = max(right, lastSmallerIndex, i)
		if i == right {
			ret++
		}
	}
	return ret
}

func main() {
	fmt.Println(maxChunksToSorted([]int{4, 2, 2, 1, 1}))
	fmt.Println(maxChunksToSorted([]int{2, 1, 3, 4, 4}))
	fmt.Println(maxChunksToSorted([]int{3}))
}
