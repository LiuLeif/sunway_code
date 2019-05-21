// 2018-11-01 11:05
package main

import (
	"fmt"
	"math"
	"sort"
)

type Pair struct {
	first, second int
}

func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {
	v := make([]Pair, len(difficulty))
	for i := 0; i < len(difficulty); i++ {
		v[i] = Pair{difficulty[i], profit[i]}
	}
	sort.Slice(v, func(i, j int) bool {
		return v[i].first < v[j].first
	})

	max := func(nums ...int) int {
		ret := math.MinInt32
		for _, n := range nums {
			if ret < n {
				ret = n
			}
		}
		return ret
	}

	maxValue := make([]int, len(v))
	for i := 0; i < len(v); i++ {
		if i == 0 {
			maxValue[i] = max(v[i].second)
		} else {
			maxValue[i] = max(maxValue[i-1], v[i].second)
		}
	}

	ceil := func(target int) int {
		lo, hi := 0, len(v)-1
		for lo <= hi {
			mid := (lo + hi) / 2
			if v[mid].first <= target {
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		}
		return hi
	}

	ret := 0
	for i := 0; i < len(worker); i++ {
		index := ceil(worker[i])
		if index != -1 {
			ret += maxValue[index]
		}
	}
	return ret
}

func main() {
	// 324
	fmt.Println(maxProfitAssignment([]int{68, 35, 52, 47, 86}, []int{67, 17, 1, 81, 3}, []int{92, 10, 85, 84, 82}))
}
