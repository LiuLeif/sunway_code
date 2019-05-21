// 2017-12-29 14:47
package main

import (
	"fmt"
	"math"
	"strings"
)

func max(nums ...int) int {
	ret := math.MinInt32
	for _, n := range nums {
		if ret < n {
			ret = n
		}
	}
	return ret
}

func longestSubstring(s string, k int) int {
	if len(s) < k {
		return 0
	}
	count := [26]int{}
	for i := 0; i < len(s); i++ {
		count[s[i]-'a']++
	}
	for i := 0; i < 26; i++ {
		if count[i] >= k || count[i] == 0 {
			continue
		}
		ret := 0
		for _, v := range strings.Split(s, string(i+'a')) {
			ret = max(ret, longestSubstring(v, k))
		}
		return ret
	}
	return len(s)
}

// note that the time complexity is O(n) instead of O(n2), since you
// can only divide_and_conquer for at most 26 times

func main() {
	fmt.Println(longestSubstring("abab", 2))
}
