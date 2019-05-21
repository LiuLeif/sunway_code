// 2018-03-22 16:04
package main

import (
	"fmt"
	"sort"
)

type MyCalendar struct {
	starts []int
	ends   map[int]int
}

func Constructor() MyCalendar {
	ret := MyCalendar{}
	ret.starts = []int{}
	ret.ends = map[int]int{}
	return ret
}

func (this *MyCalendar) Book(start int, end int) bool {
	if len(this.starts) == 0 {
		this.starts = append(this.starts, start)
		this.ends[start] = end
		return true
	}
	index := sort.Search(len(this.starts), func(i int) bool {
		return this.starts[i] > start
	})
	if index == len(this.starts) {
		x := this.starts[len(this.starts)-1]
		y := this.ends[x]
		if y > start {
			return false
		}
		this.starts = append(this.starts, start)
		this.ends[start] = end
		return true
	} else {
		if index != 0 {
			x := this.starts[index-1]
			y := this.ends[x]
			if y > start {
				return false
			}
		}
		x := this.starts[index]
		if end > x {
			return false
		}
		this.ends[start] = end
		tmp := append([]int{}, this.starts[:index]...)
		tmp = append(tmp, start)
		tmp = append(tmp, this.starts[index:]...)
		this.starts = tmp
		return true
	}
}

func main() {
	calendar := Constructor()
	fmt.Println(calendar.Book(47, 50))
	fmt.Println(calendar.Book(33, 41))
	fmt.Println(calendar.Book(39, 45))
	fmt.Println(calendar.Book(33, 42))
	fmt.Println(calendar.Book(25, 32))
	fmt.Println(calendar.Book(26, 35))
	fmt.Println(calendar.Book(19, 25))
	fmt.Println(calendar.Book(3, 8))
	fmt.Println(calendar.Book(8, 13))
	fmt.Println(calendar.Book(18, 27))
}
