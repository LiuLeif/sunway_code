// 2017-11-14 18:54
package main

import "fmt"

func max(x int, y int) int {
	if x >= y {
		return x
	}
	return y
}

func lengthOfLongestSubstring(s string) int {
	begin := 0
	m := make(map[rune]int)

	ret := 0
	for i, v := range s {
		prev, found := m[v]
		if found && prev >= begin {
			ret = max(ret, i-begin)
			begin = prev + 1
		}
		m[v] = i
	}
	return max(ret, len(s)-begin)
}

func main() {
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}
