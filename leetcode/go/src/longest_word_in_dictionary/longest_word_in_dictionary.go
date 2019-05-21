// 2018-01-05 13:53
package main

import (
	"fmt"
	"sort"
)

func longestWord(words []string) string {
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})
	ret := ""
	m := map[string]bool{}
	for i := 0; i < len(words); i++ {
		if len(words[i]) == 1 {
			m[words[i]] = true
		} else {
			m[words[i]] = m[words[i][:len(words[i])-1]]
		}
		if m[words[i]] {
			if len(words[i]) > len(ret) {
				ret = words[i]
			} else if len(words[i]) == len(ret) {
				if words[i] < ret {
					ret = words[i]
				}
			}
		}
	}
	return ret
}
func main() {
	fmt.Println(longestWord([]string{"w", "wo", "wor", "worl", "world"}))
	fmt.Println(longestWord([]string{"a", "banana", "app", "appl", "ap", "apply", "apple"}))
}
