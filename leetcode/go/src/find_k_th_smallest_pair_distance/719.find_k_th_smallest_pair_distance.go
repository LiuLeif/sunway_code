// 2018-12-10 14:26
package main

import (
	"fmt"
	"sort"
)

func smallestDistancePair(nums []int, k int) int {
	// 1,1,3
	N := len(nums)
	sort.Ints(nums)
	smallerOrEqual := func(x int) int {
		//1,1,3
		ret := 0
		for i := 0; i < N; i++ {
			for j := i + 1; j < N; j++ {
				if nums[j]-nums[i] <= x {
					ret++
				} else {
					break
				}
			}
		}
		return ret
	}
	lo, hi := 0, nums[len(nums)-1]-nums[0]
	for lo <= hi {
		mid := (lo + hi) / 2
		n := smallerOrEqual(mid)
		if n < k {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return lo
}

func main() {
	fmt.Println(smallestDistancePair([]int{1, 3, 1}, 1))
	fmt.Println(smallestDistancePair([]int{1, 6, 1}, 3))
}
