// 2018-11-26 16:29
package main

import (
	"fmt"
	"math"
	"sort"
)

func minIncrementForUnique(A []int) int {
	// 1 1 2 2 3 4 7
	// 1 2 3 4 5 6 7
	if len(A) <= 1 {
		return 0
	}
	max := func(nums ...int) int {
		ret := math.MinInt32
		for _, n := range nums {
			if ret < n {
				ret = n
			}
		}
		return ret
	}

	sort.Ints(A)
	x := A[0] + 1
	ret := 0
	for _, n := range A[1:] {
		if n < x {
			ret += x - n
			x++
		} else {
			x = max(x+1, n+1)
		}
	}
	return ret
}

func main() {
	fmt.Println(minIncrementForUnique([]int{0, 2, 2}))
}
