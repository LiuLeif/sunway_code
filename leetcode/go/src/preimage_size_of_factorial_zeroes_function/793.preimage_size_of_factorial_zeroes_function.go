// 2018-12-17 16:28
package main

import "fmt"

func count(x int) int {
	ret := 0
	k := 5
	for x >= k {
		ret += x / k
		k *= 5
	}
	return ret
}

func preimageSizeFZF(K int) int {
	lo, hi := 0, K*5
	for lo <= hi {
		mid := (lo + hi) / 2
		if count(mid) <= K {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	if count(hi) != K {
		return 0
	}
	return 5
}

func main() {
	fmt.Println(preimageSizeFZF(25))
}
