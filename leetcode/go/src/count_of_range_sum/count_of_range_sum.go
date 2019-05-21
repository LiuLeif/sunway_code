// 2018-02-06 20:11
package main

import (
	"fmt"
	"sort"
)

func countRangeSum(nums []int, lower int, upper int) int {
	sums := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		sums[i+1] = sums[i] + nums[i]
	}
	// sum of nums[i:j] = sums[j]-sums[i]
	var merge func(lo, hi int) int
	merge = func(lo, hi int) int {
		if lo == hi {
			return 0
		}
		mid := (lo + hi + 1) / 2
		ret := merge(lo, mid-1) + merge(mid, hi)
		i, j := mid, mid
		for x := lo; x < mid; x++ {
			for i <= hi && sums[i]-sums[x] < lower {
				i++
			}
			for j <= hi && sums[j]-sums[x] <= upper {
				j++
			}
			ret += j - i
		}
		sort.Ints(sums[lo : hi+1])
		return ret
	}

	return merge(0, len(sums)-1)
}
func main() {
	// [1 -3 -3 1 1 2]
	// [0 1 -2 -5 -4 -3 -1]
	fmt.Println(countRangeSum([]int{1, -3, -3, 1, 1, 2}, 3, 5))
	fmt.Println(countRangeSum([]int{-2, 2, 1}, -2, 2))
}
