// 2017-11-14 18:54
package main

import "fmt"

func dfs(s string, ret *[][]string, curr *[]string, dp [][]bool, start int) {
	n := len(s)
	if start == n {
		t := make([]string, len(*curr))
		copy(t, *curr)
		*ret = append(*ret, t)
		return
	}
	for i := start; i < n; i++ {
		if dp[start][i] {
			*curr = append(*curr, s[start:i+1])
			dfs(s, ret, curr, dp, i+1)
			*curr = (*curr)[:len(*curr)-1]
		}
	}
}

func partition(s string) [][]string {
	ret := make([][]string, 0)
	n := len(s)
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}

	for i := 0; i < n; i++ {
		dp[i][i] = true
	}

	for i := 1; i < n; i++ {
		for j := 0; j < n-i; j++ {
			if i == 1 {
				if s[j] == s[j+i] {
					dp[j][i+j] = true
				} else {
					dp[j][i+j] = false
				}
			} else {
				if s[j] == s[j+i] {
					dp[j][i+j] = dp[j+1][i+j-1]
				} else {
					dp[j][i+j] = false
				}
			}
		}
	}

	dfs(s, &ret, &[]string{}, dp, 0)
	return ret
}
func main() {
	fmt.Println(partition("aaa"))
}
