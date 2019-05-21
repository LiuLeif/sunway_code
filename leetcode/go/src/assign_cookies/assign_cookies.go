// 2017-12-28 14:52
package main

import (
	"fmt"
	"sort"
)

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	a, b := 0, 0
	for a < len(g) && b < len(s) {
		if g[a] <= s[b] {
			a++
			b++
		} else {
			b++
		}
	}
	return a
}
func main() {
	fmt.Println(findContentChildren([]int{1, 2, 3}, []int{2, 1, 3}))
}
