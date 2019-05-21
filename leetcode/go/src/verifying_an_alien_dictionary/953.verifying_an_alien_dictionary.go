// 2018-12-11 11:18
package main

import (
	"fmt"
	"sort"
)

func isAlienSorted(words []string, order string) bool {
	m := map[byte]byte{}
	for i, c := range order {
		m[byte(c)] = byte('a' + i)
	}
	return sort.SliceIsSorted(words, func(i, j int) bool {
		s1 := []byte(words[i])
		s2 := []byte(words[j])
		for i := 0; i < len(s1); i++ {
			s1[i] = m[s1[i]]
		}
		for i := 0; i < len(s2); i++ {
			s2[i] = m[s2[i]]
		}
		return string(s1) < string(s2)
	})
}

func main() {
	fmt.Println(isAlienSorted([]string{"hello", "leetcode"}, "hlabcdefgijkmnopqrstuvwxyz"))
	fmt.Println(isAlienSorted([]string{"apple", "app"}, "abcdefghijklmnopqrstuvwxyz"))
}
