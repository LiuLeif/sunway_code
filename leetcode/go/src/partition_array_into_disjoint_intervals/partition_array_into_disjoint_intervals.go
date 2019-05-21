// 2018-10-17 17:23
package main

import "fmt"

func partitionDisjoint2(A []int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	maxVal := A[0]
	cachedMaxVal := A[0]
	ret := 0
	mode := 0
	for i := 1; i < len(A); i++ {
		if mode == 0 {
			if A[i] >= maxVal {
				ret = i - 1
				mode = 1
				cachedMaxVal = maxVal
			}
		} else {
			if A[i] < cachedMaxVal {
				mode = 0
			}
		}
		maxVal = max(maxVal, A[i])
	}
	return ret + 1
}

func partitionDisjoint(A []int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}

	maxLeft := make([]int, len(A))
	maxLeft[0] = A[0]
	for i := 1; i < len(A); i++ {
		maxLeft[i] = max(maxLeft[i-1], A[i])
	}

	minRight := make([]int, len(A))
	minRight[len(minRight)-1] = A[len(A)-1]
	for i := len(minRight) - 1; i > 0; i-- {
		minRight[i-1] = min(minRight[i], A[i])
	}

	for i := 0; i < len(A); i++ {
		if maxLeft[i] <= minRight[i] {
			return i + 1
		}
	}
	return 0
}

func main() {
	fmt.Println(partitionDisjoint([]int{1, 1}))
}
