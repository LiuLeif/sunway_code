// 2018-11-09 10:22
package main

import (
	"fmt"
	"sort"
)

type MyCalendarTwo struct {
	counter map[int]int
}

func Constructor() MyCalendarTwo {
	return MyCalendarTwo{map[int]int{}}
}

func (this *MyCalendarTwo) Book(start int, end int) bool {
	this.counter[start]++
	this.counter[end]--
	keys := []int{}
	for k, _ := range this.counter {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	total := 0
	for _, k := range keys {
		total += this.counter[k]
		if total > 2 {
			this.counter[start]--
			this.counter[end]++
			return false
		}
	}
	return true
}

func main() {
	calendar := Constructor()
	fmt.Println(calendar.Book(10, 20)) // returns true
	fmt.Println(calendar.Book(50, 60)) // returns true
	fmt.Println(calendar.Book(10, 40)) // returns true
	fmt.Println(calendar.Book(5, 15))  // returns false
	fmt.Println(calendar.Book(5, 10))  // returns true
	fmt.Println(calendar.Book(25, 55)) // returns true
}
