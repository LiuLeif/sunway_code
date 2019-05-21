// 2018-02-09 13:42
package main

import "fmt"

func helper(m, n, N, i, j int, mem map[[3]int]int) int {
	if N == 0 || i < 0 || i >= m || j < 0 || j >= n {
		return 0
	}
	if i-N >= 0 && i+N < m && j-N >= 0 && j+N < n {
		return 0
	}
	if x, ok := mem[[3]int{i, j, N}]; ok {
		return x
	}
	ret := 0
	if i == 0 {
		ret++
	}
	if i == m-1 {
		ret++
	}
	if j == 0 {
		ret++
	}
	if j == n-1 {
		ret++
	}
	ret += findPaths(m, n, N-1, i-1, j)%1000000007 +
		findPaths(m, n, N-1, i+1, j)%1000000007 +
		findPaths(m, n, N-1, i, j-1)%1000000007 +
		findPaths(m, n, N-1, i, j+1)%1000000007
	ret %= 1000000007
	mem[[3]int{i, j, N}] = ret
	return ret
}

func findPathsRecursive(m int, n int, N int, i int, j int) int {
	mem := map[[3]int]int{}
	return helper(m, n, N, i, j, mem)
}

func findPaths(m int, n int, N int, i int, j int) int {
	dp := make([][][]int, N+1)
	for i := 0; i < N+1; i++ {
		dp[i] = make([][]int, m)
		for j := 0; j < m; j++ {
			dp[i][j] = make([]int, n)
		}
	}
	paths := func(i, j, k int) int {
		if j < 0 || j >= m || k < 0 || k >= n {
			return 1
		} else {
			return dp[i][j][k]
		}
	}
	for i := 1; i < N+1; i++ {
		for j := 0; j < m; j++ {
			for k := 0; k < n; k++ {
				dp[i][j][k] = paths(i-1, j-1, k) +
					paths(i-1, j+1, k) +
					paths(i-1, j, k-1) +
					paths(i-1, j, k+1)
				dp[i][j][k] %= 1000000007
			}
		}
	}
	return dp[N][i][j]
}
func main() {
	// fmt.Println(findPaths(1, 3, 3, 0, 1))
	// fmt.Println(findPaths(10, 10, 11, 5, 5))
	fmt.Println(findPaths(8, 50, 23, 5, 26))
}
