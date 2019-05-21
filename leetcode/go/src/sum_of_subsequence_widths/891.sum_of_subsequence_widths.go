// 2018-11-02 12:28
package main

import (
	"fmt"
	"math"
	"sort"
)

var MOD = int(math.Pow(10, 9)) + 7

func powOf(k int) func(x int) int {
	cache := map[int]int{}
	f := func(x int) int {
		if tmp, ok := cache[x]; ok {
			return tmp
		}
		ret := 1
		for i := 0; i < x; i++ {
			ret = k * ret
			ret %= MOD
			cache[i+1] = ret
		}
		cache[x] = ret
		return ret
	}
	return f
}

func sumSubseqWidths(A []int) int {
	sort.Ints(A)
	pow := powOf(2)
	ret := 0
	for i := 0; i < len(A); i++ {
		ret += pow(i) * A[i]
		ret -= pow(len(A)-i-1) * A[i]
		ret %= MOD
	}
	return ret
}

func main() {
	// 1 0 2 = 4+1+12=17
	// 1 2 0 = 4+4+3 = 11
	fmt.Println(sumSubseqWidths([]int{3, 7, 2, 3}))
	fmt.Println(sumSubseqWidths([]int{96, 87, 191, 197, 40, 101, 108, 35, 169, 50, 168, 182, 95, 80, 144, 43, 18, 60, 174, 13, 77, 173, 38, 46, 80, 117, 13, 19, 11, 6, 13, 118, 39, 80, 171, 36, 86, 156, 165, 190, 53, 49, 160, 192, 57, 42, 97, 35, 124, 200, 84, 70, 145, 180, 54, 141, 159, 42, 66, 66, 25, 95, 24, 136, 140, 159, 71, 131, 5, 140, 115, 76, 151, 137, 63, 47, 69, 164, 60, 172, 153, 183, 6, 70, 40, 168, 133, 45, 116, 188, 20, 52, 70, 156, 44, 27, 124, 59, 42, 172}))
}
