// 2018-10-19 13:51
package main

import (
	"fmt"
	"sort"
)

type ExamRoom struct {
	N     int
	count int
	mark  map[int]bool
}

func Constructor(N int) ExamRoom {
	ret := ExamRoom{N: N}
	ret.mark = map[int]bool{}
	return ret
}

func (this *ExamRoom) Seat() int {
	if this.count == 0 {
		this.count++
		this.mark[0] = true
		return 0
	}

	prevPos := -1
	maxLen := 0
	ret := 0
	keys := []int{}
	for k, v := range this.mark {
		if v {
			keys = append(keys, k)
		}
	}
	sort.Ints(keys)
	for _, i := range keys {
		if this.mark[i] == true {
			if prevPos == -1 {
				maxLen = i
				ret = 0
			} else {
				// 0 1 2 3 4
				len := (i - prevPos) / 2
				if len > maxLen {
					maxLen = len
					ret = (i + prevPos) / 2
				}
			}
			prevPos = i
		}
	}

	len := this.N - prevPos - 1
	if len > maxLen {
		maxLen = len
		ret = this.N - 1
	}
	this.count++
	this.mark[ret] = true
	return ret
}

func (this *ExamRoom) Leave(p int) {
	this.mark[p] = false
	this.count--
}

func main() {
	room := Constructor(10)
	fmt.Println(room.Seat())
	fmt.Println(room.Seat())
	fmt.Println(room.Seat())
}
