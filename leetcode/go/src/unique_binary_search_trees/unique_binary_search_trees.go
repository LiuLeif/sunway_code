// 2018-01-17 17:50
package main

import "fmt"

func numTrees(n int) int {
	if n < 3 {
		return n
	}
	dp := make([]int, n+1)
	dp[0], dp[1], dp[2] = 1, 1, 2

	for i := 3; i < n+1; i++ {
		for j := 0; j < i; j++ {
			dp[i] += dp[j] * dp[i-j-1]
		}
	}
	return dp[n]
}
func main() {
	// 1 2 3
	fmt.Println(numTrees(3))
}
