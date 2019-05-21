// 2018-11-16 21:47
package main

import (
	"fmt"
	"math"
)

var MOD = int(math.Pow(10, 9)) + 7

func knightDialer(N int) int {
	// 1 2 3
	// 4 5 6
	// 7 8 9
	//   0
	next := [10][]int{}
	next[0] = []int{4, 6}
	next[1] = []int{6, 8}
	next[2] = []int{7, 9}
	next[3] = []int{4, 8}
	next[4] = []int{0, 3, 9}
	next[5] = []int{}
	next[6] = []int{0, 1, 7}
	next[7] = []int{2, 6}
	next[8] = []int{1, 3}
	next[9] = []int{2, 4}
	dp := [10]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}

	for i := 1; i < N; i++ {
		dpNew := [10]int{}
		for j := 0; j < 10; j++ {
			for _, neighbor := range next[j] {
				dpNew[j] += dp[neighbor]
				dpNew[j] %= MOD
			}
		}
		dp = dpNew
	}

	ret := 0
	for i := 0; i < 10; i++ {
		ret += dp[i]
		ret %= MOD
	}
	return ret
}

func main() {
	fmt.Println(knightDialer(161)) //
	fmt.Println(knightDialer(2))   // 20
	fmt.Println(knightDialer(3))   // 46
}
