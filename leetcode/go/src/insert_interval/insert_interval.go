// 2018-01-30 16:01
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

func insert(intervals []Interval, newInterval Interval) []Interval {
	ret := []Interval{}
	intervals = append(intervals, newInterval)
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i].Start < intervals[j].Start {
			return true
		}
		return false
	})

	left, right := intervals[0].Start, intervals[0].End
	for i := 1; i < len(intervals); i++ {
		interval := intervals[i]
		if interval.Start > right {
			ret = append(ret, Interval{left, right})
			left = interval.Start
			right = interval.End
		} else {
			right = max(right, interval.End)
		}
	}
	ret = append(ret, Interval{left, right})
	return ret
}
func main() {
	fmt.Println(insert([]Interval{
		{1, 3},
		{6, 9},
	}, Interval{2, 5}))
}
