// 2018-02-24 16:02
package main

import "fmt"

func sum(nums ...int) int {
	ret := 0
	for _, n := range nums {
		ret += n
	}
	return ret
}

func countNumbersWithUniqueDigits(n int) int {
	if n == 0 {
		return 1
	}
	// [0,100)
	// 11,22,33,44,55...99
	dp := make([]int, n+1)
	dp[1] = 9
	// 1 2 3 4 5 6 7 8 9
	//
	for i := 2; i < n+1; i++ {
		if i >= 11 {
			break
		}
		dp[i] = dp[i-1] * (11 - i)
	}
	return sum(dp...) + 1
}

func countNumbersWithUniqueDigits2(n int) int {
	if n == 0 {
		return 1
	}
	ret := 10
	avail := 9
	curr := 9
	for n > 1 && avail > 0 {
		curr *= avail
		ret += curr
		n--
		avail--
	}
	return ret
}

func main() {
	fmt.Println(countNumbersWithUniqueDigits(13))
	fmt.Println(countNumbersWithUniqueDigits2(13))
}
