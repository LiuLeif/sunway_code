// 2018-03-30 14:53
package main

import "fmt"

func champagneTower(poured int, query_row int, query_glass int) float64 {
	dp := make([][]float64, query_row+1)
	for i := 0; i < query_row+1; i++ {
		dp[i] = make([]float64, query_row+1)
	}

	get := func(i, j int) float64 {
		if j < 0 {
			return 0
		}
		if j > i {
			return 0
		}
		if dp[i][j] <= 1 {
			return 0
		}
		return (dp[i][j] - 1) / 2.0
	}

	dp[0][0] = float64(poured)
	for i := 1; i < query_row+1; i++ {
		cont := false
		for j := 0; j <= i; j++ {
			dp[i][j] = get(i-1, j-1) + get(i-1, j)
			if dp[i][j] != 0 {
				cont = true
			}
		}
		if !cont {
			break
		}
	}
	if dp[query_row][query_glass] > 1 {
		return 1
	}
	return dp[query_row][query_glass]
}
func main() {
	fmt.Println(champagneTower(2, 1, 0))
}
