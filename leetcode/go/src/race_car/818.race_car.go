// 2018-12-13 10:41
package main

import (
	"fmt"
	"strings"
)

type Tuple struct {
	pos, speed, count int
	action            string
}

type Queue []Tuple

func (this *Queue) Poll() Tuple {
	tmp := (*this)[0]
	*this = (*this)[1:]
	return tmp
}

func (this *Queue) Offer(x Tuple) {
	*this = append(*this, x)
}

func (this *Queue) Empty() bool {
	return len(*this) == 0
}

func (this *Queue) Peek() Tuple {
	return (*this)[0]
}

func racecar(target int) int {
	// 0 1 3 3 2 2 3 5
	queue := Queue{}
	queue.Offer(Tuple{0, 1, 0, ""})

	visited := map[[2]int]bool{}
	visited[[2]int{0, 1}] = true

	for !queue.Empty() {
		top := queue.Poll()
		if top.pos == target {
			return top.count
		}
		pos, speed := top.pos+top.speed, top.speed*2
		if !visited[[2]int{pos, speed}] && (pos > 0 && pos < 20000) {
			visited[[2]int{pos, speed}] = true
			queue.Offer(Tuple{pos, speed, top.count + 1, top.action + "a"})
		}
		pos = top.pos
		if top.speed > 0 {
			speed = -1
		} else {
			speed = 1
		}
		if !visited[[2]int{pos, speed}] && !strings.HasSuffix(top.action, "rr") && (pos > 0 && pos < 20000) {
			visited[[2]int{pos, speed}] = true
			queue.Offer(Tuple{pos, speed, top.count + 1, top.action + "r"})
		}
	}
	return 0
}

func main() {
	fmt.Println(racecar(5))
	fmt.Println(racecar(4))
	fmt.Println(racecar(6))
	fmt.Println(racecar(3))
}
