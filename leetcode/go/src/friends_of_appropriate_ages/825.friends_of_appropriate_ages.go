// 2018-11-08 15:03
package main

import (
	"fmt"
	"sort"
)

func numFriendRequests(ages []int) int {
	// 16 16 16
	sort.Slice(ages, func(i, j int) bool {
		return ages[i] > ages[j]
	})
	ret := 0
	counter := map[int]int{}
	for i := 0; i < len(ages); i++ {
		counter[ages[i]]++
	}
	for k, v := range counter {
		if k > 14 {
			ret += v * (v - 1) / 2
		}
	}
	lo, hi := 0, 0
	N := len(ages)
	for ; lo < N; lo++ {
		for ; hi < N; hi++ {
			if float64(ages[hi]) <= 0.5*float64(ages[lo])+7 {
				break
			}
		}
		if hi-1-lo > 0 {
			ret += hi - 1 - lo
		}
	}
	return ret
}

func main() {
	fmt.Println(numFriendRequests([]int{108, 115, 5, 24, 82}))
}
