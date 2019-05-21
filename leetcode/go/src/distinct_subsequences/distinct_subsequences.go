// 2018-01-31 16:19
package main

import "fmt"

func helper(s, t string, m map[[2]int]int, x, y int) int {
	if y == len(t) {
		return 1
	}
	if x == len(s) {
		return 0
	}
	if _, ok := m[[2]int{x, y}]; ok {
		return m[[2]int{x, y}]
	}
	ret := 0
	if s[x] == t[y] {
		ret = helper(s, t, m, x+1, y+1) + helper(s, t, m, x+1, y)
	} else {
		ret = helper(s, t, m, x+1, y)
	}
	m[[2]int{x, y}] = ret
	return ret
}

func numDistinct(s string, t string) int {
	// rabbbit
	// rabbit
	m := map[[2]int]int{}
	return helper(s, t, m, 0, 0)
}
func main() {
	fmt.Println(numDistinct("rabbbit", "rabbit"))
	fmt.Println(numDistinct("abbabc", "bc"))
}
