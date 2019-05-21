// 2018-02-26 14:41
package main

import "fmt"

type Tuple struct {
	i, x, y int
}

func knightProbability(N int, K int, r int, c int) float64 {
	dp := map[Tuple]float64{}

	get := func(i, x, y int) float64 {
		if x < 0 || x >= N || y < 0 || y >= N {
			return 0.0
		}
		if i < 0 {
			return 1
		}
		return dp[Tuple{i, x, y}]
	}
	step := []int{-2, -1, 1, 2}
	for i := 0; i < K; i++ {
		for x := 0; x < N; x++ {
			for y := 0; y < N; y++ {
				for k := 0; k < len(step); k++ {
					for l := 0; l < len(step); l++ {
						if step[k] == step[l] || step[k] == -step[l] {
							continue
						}
						dp[Tuple{i, x, y}] += get(i-1, x+step[k], y+step[l]) / 8.0
					}
				}
			}
		}
	}
	return get(K-1, r, c)
}
func main() {
	fmt.Println(knightProbability(3, 2, 0, 0))
}
