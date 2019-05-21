// 2017-11-30 10:51
package main

import "fmt"
import "strconv"

func numDecodings(s string) int {
	if len(s) == 0 {
		return 0
	}
	dp := make([]int, len(s)+1)
	dp[len(s)] = 1
	if s[len(s)-1] == '0' {
		dp[len(s)-1] = 0
	} else {
		dp[len(s)-1] = 1
	}
	for i := len(s) - 2; i >= 0; i-- {
		if s[i] == '0' {
			dp[i] = 0
			continue
		}
		x, _ := strconv.Atoi(s[i : i+2])
		if x <= 26 {
			dp[i] = dp[i+1] + dp[i+2]
		} else {
			dp[i] = dp[i+1]
		}
	}

	return dp[0]
}
func main() {
	fmt.Println(numDecodings("01"))
}
