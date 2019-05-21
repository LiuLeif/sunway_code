// 2018-10-26 15:19
package main

import (
	"fmt"
	"sort"
)

func advantageCount(A []int, B []int) []int {
	ret := make([]int, len(A))
	sort.Ints(A)

	firstLarger := func(target int) int {
		lo, hi := 0, len(A)-1
		for lo <= hi {
			mid := (lo + hi) / 2
			if A[mid] <= target {
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		}
		if lo == len(A) {
			lo = 0
		}
		ret := A[lo]
		A = append(A[:lo], A[lo+1:]...)
		return ret
	}
	for i := 0; i < len(B); i++ {
		ret[i] = firstLarger(B[i])
	}
	return ret
}

func main() {
	fmt.Println(advantageCount([]int{12, 24, 8, 32}, []int{13, 25, 32, 11}))
}
