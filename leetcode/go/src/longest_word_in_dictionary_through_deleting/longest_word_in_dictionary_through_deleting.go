// 2018-01-11 17:37
package main

import (
	"fmt"
	"sort"
)

func findLongestWord(s string, d []string) string {
	sort.Slice(d, func(i, j int) bool {
		switch {
		case len(d[i]) > len(d[j]):
			return true
		case len(d[i]) < len(d[j]):
			return false
		default:
			return d[i] < d[j]
		}
	})

	for i := 0; i < len(d); i++ {
		w := d[i]
		p1, p2 := 0, 0
		for p1 < len(s) && p2 < len(w) {
			if s[p1] == w[p2] {
				p1++
				p2++
			} else {
				p1++
			}
		}
		if p2 == len(w) {
			return w
		}
	}

	return ""
}
func main() {
	fmt.Println(findLongestWord("abpcple", []string{"ale", "apple", "monkey", "plea"}))
}
