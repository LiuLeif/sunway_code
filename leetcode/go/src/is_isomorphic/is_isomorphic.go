// 2017-11-14 18:54
package main

import "fmt"

func isIsomorphic(s string, t string) bool {
	m1 := make(map[byte]byte)
	m2 := make(map[byte]byte)
	for i := 0; i < len(s); i++ {
		x := s[i]
		y := t[i]
		if m1[x] != 0 && m1[x] != y {
			return false
		}
		if m2[y] != 0 && m2[y] != x {
			return false
		}
		m1[x] = y
		m2[y] = x
	}

	return true
}
func main() {
	fmt.Println(isIsomorphic("ab", "ca"))
}
