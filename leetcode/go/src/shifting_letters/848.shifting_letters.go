// 2018-10-22 15:16
package main

import "fmt"

func shiftingLetters(S string, shifts []int) string {
	dp := make([]int, len(S))
	dp[len(dp)-1] = shifts[len(dp)-1]
	for i := len(S) - 1; i >= 1; i-- {
		dp[i-1] = shifts[i-1] + dp[i]
	}
	chars := []rune(S)
	for i := 0; i < len(chars); i++ {
		chars[i] = rune((int(chars[i])-'a'+dp[i])%26 + 'a')
	}
	return string(chars)
}

func main() {
	fmt.Println(shiftingLetters("abc", []int{3, 5, 9}))
}
