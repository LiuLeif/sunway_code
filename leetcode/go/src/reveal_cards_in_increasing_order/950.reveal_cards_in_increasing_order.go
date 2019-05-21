// 2018-12-03 10:05
package main

import (
	"fmt"
	"sort"
)

type Queue []int

func (this *Queue) Poll() int {
	tmp := (*this)[0]
	*this = (*this)[1:]
	return tmp
}

func (this *Queue) Offer(x int) {
	*this = append(*this, x)
}

func (this *Queue) Empty() bool {
	return len(*this) == 0
}

func (this *Queue) Peek() int {
	return (*this)[0]
}

func deckRevealedIncreasing(deck []int) []int {
	sort.Ints(deck)
	N := len(deck)
	ret := make([]int, N)
	queue := Queue{}
	for i := 0; i < N; i++ {
		queue.Offer(i)
	}

	for i := 0; i < N; i++ {
		ret[queue.Poll()] = deck[i]
		if !queue.Empty() {
			queue.Offer(queue.Poll())
		}
	}

	return ret
}

func main() {
	// 2 3 5 7 11 13 17
	// 2 3 5 7 11 13 17 3 7 13 17 3
	fmt.Println(deckRevealedIncreasing([]int{17, 13, 11, 2, 3, 5, 7}))
}
