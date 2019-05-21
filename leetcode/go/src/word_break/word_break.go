// 2017-12-13 13:52
package main

import (
	"fmt"
	"strings"
)

func dfs(visited map[string]bool, s string, wordDict []string) bool {
	if len(s) == 0 {
		return true
	}
	for _, word := range wordDict {
		if strings.HasPrefix(s, word) {
			if visited[s[len(word):]] {
				continue
			}
			visited[s[len(word):]] = true
			ret := dfs(visited, s[len(word):], wordDict)
			if ret {
				return true
			}
		}
	}
	return false
}
func wordBreak(s string, wordDict []string) bool {
	visited := make(map[string]bool)
	return dfs(visited, s, wordDict)
}
func wordBreakDP(s string, wordDict []string) bool {
	m := make(map[string]bool)
	for _, word := range wordDict {
		m[word] = true
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 0; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && m[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}
func main() {
	s := "leetcode"
	dict := []string{"leet", "code"}
	fmt.Println(wordBreakDP(s, dict))
}
