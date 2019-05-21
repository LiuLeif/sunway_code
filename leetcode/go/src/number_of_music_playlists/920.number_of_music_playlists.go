// 2018-12-17 10:35
package main

import (
	"fmt"
	"math"
)

var MOD = int(math.Pow(10, 9)) + 7

func numMusicPlaylists(N int, L int, K int) int {
	dp := make([][]int, L+1)
	for i := 0; i < L+1; i++ {
		dp[i] = make([]int, N+1)
	}
	dp[0][0] = 1
	for i := 1; i < L+1; i++ {
		for j := 1; j < N+1; j++ {
			dp[i][j] = dp[i-1][j-1] * (N - (j - 1)) % MOD
			if j > K {
				dp[i][j] = dp[i][j] + dp[i-1][j]*(j-K)%MOD
			}
		}
	}
	return dp[L][N] % MOD
}

func main() {
	fmt.Println(numMusicPlaylists(2, 3, 1))
	fmt.Println(numMusicPlaylists(25, 28, 5))
}
