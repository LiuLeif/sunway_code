// 2017-11-14 18:54
package main

import "fmt"

func longestPalindrome(s string) int {
	m := make(map[rune]int)
	for _, v := range s {
		m[v] += 1
	}
	ret := 0
	extra := 0
	for _, v := range m {
		ret += (v / 2) * 2
		if v%2 != 0 {
			extra = 1
		}
	}
	return ret + extra
}
func main() {
	fmt.Println(longestPalindrome("a"))
}
