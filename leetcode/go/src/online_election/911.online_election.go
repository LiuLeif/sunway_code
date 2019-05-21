// 2018-10-26 16:57
package main

import "fmt"

type TopVotedCandidate struct {
	winner []int
	times  []int
}

func Constructor(persons []int, times []int) TopVotedCandidate {
	ret := TopVotedCandidate{}
	ret.winner = make([]int, len(persons))
	ret.times = times
	counter := map[int]int{}
	maxVote := 0
	maxWinner := 0
	for i := 0; i < len(persons); i++ {
		p := persons[i]
		counter[p]++
		if counter[p] >= maxVote {
			ret.winner[i] = p
			maxVote = counter[p]
			maxWinner = p
		} else {
			ret.winner[i] = maxWinner
		}
	}
	return ret
}

func (this *TopVotedCandidate) Q(t int) int {
	search := func(target int) int {
		lo, hi := 0, len(this.times)-1
		for lo <= hi {
			mid := (lo + hi) / 2
			if this.times[mid] == target {
				return mid
			} else if this.times[mid] > target {
				hi = mid - 1
			} else {
				lo = mid + 1
			}
		}
		return hi
	}

	index := search(t)
	return this.winner[index]
}

func main() {
	candidates := Constructor([]int{0, 0, 0, 0, 1}, []int{0, 6, 39, 52, 75})
	fmt.Println(candidates.Q(45))
	fmt.Println(candidates.Q(49))
	fmt.Println(candidates.Q(59))
	fmt.Println(candidates.Q(68))
	fmt.Println(candidates.Q(42))
	fmt.Println(candidates.Q(37))
	fmt.Println(candidates.Q(99))
	fmt.Println(candidates.Q(26))
	fmt.Println(candidates.Q(78))
	fmt.Println(candidates.Q(43))
}
