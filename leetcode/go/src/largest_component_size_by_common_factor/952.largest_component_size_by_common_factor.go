// 2018-12-13 15:21
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

func largestComponentSize(A []int) int {
	parent := map[int]int{}
	rank := map[int]int{}

	find := func(x int) int {
		if _, ok := parent[x]; !ok {
			parent[x] = x
		}
		for parent[x] != x {
			x = parent[x]
		}
		return x
	}

	union := func(x, y int) {
		x, y = find(x), find(y)
		if x == y {
			return
		}
		rx, ry := rank[x], rank[y]
		if rx > ry {
			parent[y] = x
		} else if rx < ry {
			parent[x] = y
		} else {
			parent[x] = y
			rank[y]++
		}
	}

	for i := 0; i < len(A); i++ {
		for j := 2; j < int(math.Sqrt(float64(A[i])))+1; j++ {
			if A[i]%j == 0 {
				union(A[i], j)
				union(A[i], A[i]/j)
			}
		}
	}
	ret := 0
	counter := map[int]int{}
	for i := 0; i < len(A); i++ {
		x := find(A[i])
		counter[x]++
		ret = max(ret, counter[x])
	}
	return ret
}

func main() {
	// 20-50
	// 9-63
	fmt.Println(largestComponentSize([]int{20, 50, 9, 63}))
	fmt.Println(largestComponentSize([]int{4, 6, 15, 35}))
}
