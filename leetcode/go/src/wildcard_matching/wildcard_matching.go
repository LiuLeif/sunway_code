// 2018-01-08 15:53
package main

import "fmt"

func isMatch(s string, p string) bool {
	dp := make([][]bool, len(s)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, len(p)+1)
	}
	dp[0][0] = true
	for i := 0; i < len(p); i++ {
		if p[i] != '*' {
			break
		}
		dp[0][i+1] = true
	}
	for i := 1; i < len(s)+1; i++ {
		for j := 1; j < len(p)+1; j++ {
			if s[i-1] == p[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else if p[j-1] == '?' {
				dp[i][j] = dp[i-1][j-1]
			} else if p[j-1] == '*' {
				dp[i][j] = dp[i][j-1] || dp[i-1][j-1] || dp[i-1][j]
			}
		}
	}
	return dp[len(s)][len(p)]
}
func main() {
	fmt.Println(isMatch("aa", "*"))
	fmt.Println(isMatch("", "*"))
}
