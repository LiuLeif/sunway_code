// 2017-12-28 14:52
package main

import (
	"fmt"
	"sort"
)

type Interval struct {
	Start int
	End   int
}

func eraseOverlapIntervals(intervals []Interval) int {
	if len(intervals) == 0 {
		return 0
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].End < intervals[j].End
	})
	end := intervals[0].End
	count := 1
	for i := 1; i < len(intervals); i++ {
		if intervals[i].Start < end {
			continue
		}
		count++
		end = intervals[i].End
	}
	return len(intervals) - count
}

func main() {
	fmt.Println(eraseOverlapIntervals([]Interval{
		{1, 2},
		{2, 3},
		{3, 4},
		{1, 3},
	}))
}
