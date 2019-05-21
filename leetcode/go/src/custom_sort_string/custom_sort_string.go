// 2018-03-01 14:47
package main

import (
	"fmt"
	"sort"
)

func customSortString(S string, T string) string {
	m := map[byte]int{}
	for i := 0; i < len(S); i++ {
		m[S[i]] = i
	}
	b := []byte(T)
	sort.Slice(b, func(i, j int) bool {
		return m[b[i]] < m[b[j]]
	})
	return string(b)
}
func main() {
	fmt.Println(customSortString("cba", "abcdcabc"))
}
