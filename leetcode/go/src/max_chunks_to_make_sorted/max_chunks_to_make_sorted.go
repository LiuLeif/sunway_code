// 2018-03-01 16:33
package main

import (
	"fmt"
	"math"
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

func maxChunksToSorted(arr []int) int {
	// 0 1 2 3
	m := map[int]int{}
	for i := 0; i < len(arr); i++ {
		m[arr[i]] = i
	}
	ret := 0
	right := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] != i {
			right = max(right, m[i])
		}
		if i >= right {
			ret++
		}
	}
	return ret
}
func main() {
	fmt.Println(maxChunksToSorted([]int{0, 1, 2, 3, 4}))
	fmt.Println(maxChunksToSorted([]int{1, 0, 3, 2, 4}))
	fmt.Println(maxChunksToSorted([]int{0, 3, 2, 1, 4}))
	fmt.Println(maxChunksToSorted([]int{2, 0, 1}))
}
