// 2018-01-05 11:30
package main

import "fmt"

func canWinNimDP(n int) bool {
	// TLE
	dp := make([]bool, n+1)
	dp[0] = false
	for i := 1; i < n+1; i++ {
		if i >= 1 && dp[i-1] == false {
			dp[i] = true
		} else if i >= 2 && dp[i-2] == false {
			dp[i] = true
		} else if i >= 3 && dp[i-3] == false {
			dp[i] = true
		} else {
			dp[i] = false
		}
	}
	return dp[n]
}

func canWinNim(n int) bool {
	return n%4 != 0
}

func main() {
	// 0,1,2,3,4,5,6,7,8,9,...
	// f,t,t,t,f,t,t,t,f,t,...
	fmt.Println(canWinNim(4))
}
