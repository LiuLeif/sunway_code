// 2017-12-26 16:09
package main

import "fmt"

func countSubstrings(s string) int {
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}
	count := 0
	for i := 0; i < len(s); i++ {
		dp[i][i] = true
		count++
	}
	for i := 1; i < len(s); i++ {
		for j := 0; j < len(s)-i; j++ {
			if s[j] == s[i+j] {
				if i == 1 {
					dp[j][i+j] = true
				} else {
					dp[j][i+j] = dp[j+1][i+j-1]
				}
				if dp[j][i+j] == true {
					count++
				}
			}
		}
	}
	return count
}
func main() {
	fmt.Println(countSubstrings("fdsklf"))
}
