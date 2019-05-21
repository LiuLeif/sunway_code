// 2018-02-27 17:55
package main

import (
	"fmt"
	"sort"
)

type Interval struct {
	Start int
	End   int
}

type SummaryRanges struct {
	intervals []Interval
}

func Constructor() SummaryRanges {
	ret := SummaryRanges{}
	ret.intervals = []Interval{}
	return ret
}

func (this *SummaryRanges) SearchInterval(val int) int {
	index := sort.Search(len(this.intervals), func(i int) bool {
		return this.intervals[i].Start > val
	})
	if index == 0 {
		return len(this.intervals)
	}
	if this.intervals[index-1].Start <= val && this.intervals[index-1].End >= val {
		return index - 1
	}
	return len(this.intervals)
}

func (this *SummaryRanges) Addnum(val int) {
	x := this.SearchInterval(val)
	if x != len(this.intervals) {
		return
	}

	left := this.SearchInterval(val - 1)
	right := this.SearchInterval(val + 1)
	if left == len(this.intervals) && right == len(this.intervals) {
		insert := sort.Search(len(this.intervals), func(i int) bool {
			return this.intervals[i].Start > val
		})
		tmp := append([]Interval{}, this.intervals[:insert]...)
		tmp = append(tmp, Interval{val, val})
		tmp = append(tmp, this.intervals[insert:]...)
		this.intervals = tmp
		return
	}
	if left != len(this.intervals) {
		this.intervals[left].End = val
	}
	if right != len(this.intervals) {
		this.intervals[right].Start = val
	}
	if left != len(this.intervals) && right != len(this.intervals) {
		this.intervals[left].End = this.intervals[right].End
		this.intervals = append(this.intervals[:right], this.intervals[right+1:]...)
	}
}

func (this *SummaryRanges) Getintervals() []Interval {
	return this.intervals
}

func main() {
	obj := Constructor()
	obj.Addnum(1)
	obj.Addnum(3)
	obj.Addnum(2)
	obj.Addnum(7)
	obj.Addnum(5)
	obj.Addnum(6)
	fmt.Println(obj.Getintervals())
}
