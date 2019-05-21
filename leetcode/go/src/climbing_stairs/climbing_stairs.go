// 2017-12-19 15:43
package main

import "fmt"

func climbStairs2(n int) int {
	a, b := 1, 1
	for i := 0; i < n; i++ {
		a, b = b, a+b
	}
	return a
}
func climbStairs(n int) int {
	if n <= 1 {
		return 1
	}
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i < n+1; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}
func main() {
	fmt.Println(climbStairs(10))
	fmt.Println(climbStairs2(10))
}
