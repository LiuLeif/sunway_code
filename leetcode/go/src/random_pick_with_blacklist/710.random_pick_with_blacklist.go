// 2018-11-12 17:36
package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type Pair struct {
	a, b int
}

type Solution struct {
	blacklist []Pair
	N         int
}

func Constructor(N int, blacklist []int) Solution {
	v := []Pair{}
	sort.Ints(blacklist)
	blacklist = append(blacklist, N)
	count := 0
	for i := 0; i < len(blacklist); i++ {
		count++
		v = append(v, Pair{blacklist[i], blacklist[i] + 1 - count})
	}
	return Solution{v, N}
}

func (this *Solution) Pick() int {
	// [0,N)
	// 0,[1],[2],[3],4
	// 1 2 3 5
	// 1 1 1 2
	//
	// 0 1
	firstLarger := func(x int) int {
		lo, hi := 0, len(this.blacklist)-1
		for lo <= hi {
			mid := (lo + hi) / 2
			if this.blacklist[mid].b <= x {
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		}
		return lo
	}
	index := rand.Intn(this.N - len(this.blacklist) + 1)
	index2 := firstLarger(index)
	return this.blacklist[index2].a - (this.blacklist[index2].b - index)
}

func main() {
	shuffle := Constructor(5, []int{0, 1, 2})
	fmt.Println(shuffle)
	rand.Seed(time.Now().Unix())
	for i := 0; i < 20; i++ {
		fmt.Println(shuffle.Pick())
	}
}
