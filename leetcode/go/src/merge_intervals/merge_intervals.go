// 2017-12-22 15:50
package main

import (
	"fmt"
	"math"
	"sort"
)

type Interval struct {
	Start int
	End   int
}

func max(nums ...int) int {
	ret := math.MinInt32
	for _, n := range nums {
		if ret < n {
			ret = n
		}
	}
	return ret
}

func overlap(l1, r1, l2, r2 int) bool {
	return l2 <= r1
}

func merge(intervals []Interval) []Interval {
	if len(intervals) == 0 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})
	ret := []Interval{}
	left, right := intervals[0].Start, intervals[0].End
	for i := 1; i < len(intervals); i++ {
		if overlap(left, right, intervals[i].Start, intervals[i].End) {
			right = max(right, intervals[i].End)
		} else {
			ret = append(ret, Interval{left, right})
			left, right = intervals[i].Start, intervals[i].End
		}
	}
	ret = append(ret, Interval{left, right})
	return ret
}
func main() {
	fmt.Println(merge([]Interval{
		Interval{1, 4},
		Interval{2, 3},
	}))
}
