// 2018-10-19 10:35
package main

import (
	"fmt"
)

func minEatingSpeed(piles []int, H int) int {
	lo, hi := 1, 0
	for _, v := range piles {
		if v > hi {
			hi = v
		}
	}
	cost := func(x int) int {
		ret := 0
		for _, v := range piles {
			if v%x == 0 {
				ret += v / x
			} else {
				ret += v/x + 1
			}
		}
		return ret
	}

	for lo <= hi {
		mid := (lo + hi) / 2
		hours := cost(mid)
		if hours > H {
			lo = mid + 1
		} else if hours <= H {
			hi = mid - 1
		}
	}
	return lo
}
func main() {
	fmt.Println(minEatingSpeed([]int{332484035, 524908576, 855865114, 632922376, 222257295, 690155293, 112677673, 679580077, 337406589, 290818316, 877337160, 901728858, 679284947, 688210097, 692137887, 718203285, 629455728, 941802184}, 823855818))
}
