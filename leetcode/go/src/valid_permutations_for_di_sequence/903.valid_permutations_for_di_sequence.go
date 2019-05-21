// 2018-11-23 15:54
package main

import (
	"fmt"
	"math"
)

var MOD = int(math.Pow(10, 9)) + 7

func numPermsDISequence(S string) int {
	N := len(S)
	dp := make([][]int, N+1)
	for i := 0; i < N+1; i++ {
		dp[i] = make([]int, N+1)
	}
	// D: 10
	// DD: 210
	// DI: 201, 102
	dp[0][0] = 1
	for i := 1; i < N+1; i++ {
		c := S[i-1]
		for j := 0; j <= i; j++ {
			if c == 'D' {
				for k := j; k < i; k++ {
					dp[i][j] += dp[i-1][k] % MOD
				}
			} else {
				for k := 0; k < j; k++ {
					dp[i][j] += dp[i-1][k] % MOD
				}
			}
		}
	}
	ret := 0
	for i := 0; i < N+1; i++ {
		ret += dp[N][i] % MOD
	}
	return ret % MOD
}

func main() {
	fmt.Println(numPermsDISequence("DID"))
}
