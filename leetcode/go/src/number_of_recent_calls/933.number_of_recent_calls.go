// 2018-11-06 15:51
package main

import "fmt"

type RecentCounter struct {
	pings []int
}

func Constructor() RecentCounter {
	return RecentCounter{[]int{}}
}

func (this *RecentCounter) Ping(t int) int {
	this.pings = append(this.pings, t)

	firstSmaller := func(target int) int {
		lo, hi := 0, len(this.pings)-1
		for lo <= hi {
			mid := (lo + hi) / 2
			if this.pings[mid] < target {
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		}
		return lo
	}
	return len(this.pings) - firstSmaller(t-3000)
}

func main() {
	counter := Constructor()
	fmt.Println(counter.Ping(1))
	fmt.Println(counter.Ping(100))
	fmt.Println(counter.Ping(3001))
	fmt.Println(counter.Ping(3002))
}
