// 2018-01-23 14:27
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

func radixSort(nums []int, maxVal int) {
	for exp := 1; maxVal/exp > 0; exp *= 10 {
		count := [10][]int{}
		for _, num := range nums {
			pos := (num / exp) % 10
			count[pos] = append(count[pos], num)
		}
		k := 0
		for i := 0; i < 10; i++ {
			for _, v := range count[i] {
				nums[k] = v
				k++
			}
		}
	}
}

func maximumGap(nums []int) int {
	maxVal := max(nums...)
	radixSort(nums, maxVal)
	gap := 0
	for i := 1; i < len(nums); i++ {
		gap = max(gap, nums[i]-nums[i-1])
	}
	return gap
}
func main() {
	// 0 1 3 5
	fmt.Println(maximumGap([]int{1, 1, 1, 1, 1, 5, 5, 5, 5, 5}))
}
