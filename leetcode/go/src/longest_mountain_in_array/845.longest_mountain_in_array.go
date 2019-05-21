// 2018-10-22 17:22
package main

import (
	"fmt"
	"math"
)

func longestMountain(A []int) int {
	max := func(nums ...int) int {
		ret := math.MinInt32
		for _, n := range nums {
			if ret < n {
				ret = n
			}
		}
		return ret
	}

	indexOne := -1
	indexTwo := -1
	ret := 0
	for i := 1; i < len(A); i++ {
		if A[i] > A[i-1] {
			if indexOne == -1 {
				indexOne = i - 1
			} else if indexTwo != -1 {
				//
				ret = max(ret, i-indexOne)
				indexOne = i - 1
				indexTwo = -1
			}
		} else if A[i] == A[i-1] {
			if indexTwo != -1 {
				ret = max(ret, i-indexOne)
			}
			indexOne = -1
			indexTwo = -1
		} else {
			if indexOne != -1 {
				if indexTwo == -1 {
					indexTwo = i
				}
			}
		}
	}
	if indexOne != -1 && indexTwo != -1 {
		ret = max(ret, len(A)-indexOne)
	}
	return ret
}

func main() {
	// 2, 1, 4, 7, 3, 5, 2, 3, 4, 1
	fmt.Println(longestMountain([]int{2, 1, 4, 7, 3, 2, 5}))
}
