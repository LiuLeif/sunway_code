// 2018-11-07 13:05
package main

import (
	"fmt"
	"math"
	"sort"
)

type MyCalendarThree struct {
	counter map[int]int
}

func Constructor() MyCalendarThree {
	return MyCalendarThree{map[int]int{}}
}

func (this *MyCalendarThree) Book(start int, end int) int {
	max := func(nums ...int) int {
		ret := math.MinInt32
		for _, n := range nums {
			if ret < n {
				ret = n
			}
		}
		return ret
	}

	this.counter[start]++
	this.counter[end]--
	ret := 0
	total := 0
	keys := []int{}
	for k, _ := range this.counter {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		total += this.counter[k]
		ret = max(ret, total)
	}
	return ret
}

func main() {
	calendar := Constructor()
	fmt.Println(calendar.Book(10, 20)) // returns 1
	fmt.Println(calendar.Book(50, 60)) // returns 1
	fmt.Println(calendar.Book(10, 40)) // returns 2
	fmt.Println(calendar.Book(5, 15))  // returns 3
	fmt.Println(calendar.Book(5, 10))  // returns 3
	fmt.Println(calendar.Book(25, 55)) // returns 3
}
