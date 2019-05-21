// 2018-12-04 10:18
package main

import (
	"fmt"
	"math"
	"sort"
)

func bagOfTokensScore(tokens []int, P int) int {
	if len(tokens) == 0 {
		return 0
	}
	max := func(nums ...int) int {
		ret := math.MinInt32
		for _, n := range nums {
			if ret < n {
				ret = n
			}
		}
		return ret
	}

	sort.Ints(tokens)
	lo, hi := 0, len(tokens)-1
	ret := 0
	points := 0
	for lo <= hi {
		// fmt.Println(lo, hi, P, ret)
		if P >= tokens[lo] {
			P -= tokens[lo]
			lo++
			points++
			ret = max(ret, points)
		} else if points > 0 {
			P += tokens[hi]
			hi--
			points--
		} else {
			break
		}
	}
	return ret
}

func main() {
	fmt.Println(bagOfTokensScore([]int{100, 200, 300, 400}, 300))
	fmt.Println(bagOfTokensScore([]int{100, 200}, 150))
	fmt.Println(bagOfTokensScore([]int{55, 71, 82}, 54))
}
