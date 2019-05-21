// 2018-03-22 10:23
package main

import (
	"fmt"
	"sort"
)

func numMatchingSubseq(S string, words []string) int {
	ret := 0
	pos := [26][]int{}
	for i := 0; i < len(S); i++ {
		pos[S[i]-'a'] = append(pos[S[i]-'a'], i)
	}

	isSubseq := func(w string) bool {
		if len(w) > len(S) {
			return false
		}
		current := -1
		for i := 0; i < len(w); i++ {
			c := w[i] - 'a'
			index := sort.Search(len(pos[c]), func(k int) bool {
				return pos[c][k] > current
			})
			if index == len(pos[c]) {
				return false
			}
			current = pos[c][index]
		}
		return true
	}
	for _, w := range words {
		if isSubseq(w) {
			ret++
		}
	}
	return ret
}
func main() {
	// mmmmmmmmu
	// mu,mv
	fmt.Println(numMatchingSubseq("abcde", []string{"abcde"}))
}
