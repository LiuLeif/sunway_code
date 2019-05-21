// 2018-02-26 13:41
package main

import (
	"fmt"
	"sort"
)

type Interval struct {
	Start int
	End   int
}

type Pair struct {
	index    int
	interval Interval
}

func findRightInterval(intervals []Interval) []int {
	ret := make([]int, len(intervals))
	pairs := make([]Pair, len(intervals))
	for i, interval := range intervals {
		pairs[i] = Pair{i, interval}
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].interval.Start < pairs[j].interval.Start
	})

	for i := 0; i < len(intervals); i++ {
		index := sort.Search(len(intervals), func(k int) bool {
			return pairs[k].interval.Start >= intervals[i].End
		})
		if index == len(intervals) {
			index = -1
		} else {
			index = pairs[index].index
		}
		ret[i] = index
	}
	return ret
}
func main() {
	fmt.Println(findRightInterval([]Interval{
		{1, 4},
		{2, 3},
		{3, 4},
	}))
}
