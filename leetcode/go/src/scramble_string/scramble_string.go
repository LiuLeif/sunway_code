// 2018-01-18 13:58
package main

import "fmt"

type Pair struct {
	a, b string
}

func helper(m map[Pair]bool, s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	if len(s1) == 1 && s1 != s2 {
		return false
	}
	if s1 == s2 {
		return true
	}
	count := [26]int{}
	for i := 0; i < len(s1); i++ {
		count[s1[i]-'a']++
	}
	for i := 0; i < len(s1); i++ {
		count[s2[i]-'a']--
	}
	for i := 0; i < 26; i++ {
		if count[i] != 0 {
			m[Pair{s1, s2}] = false
			return false
		}
	}
	if m[Pair{s1, s2}] {
		return true
	}
	for mid := 1; mid < len(s1); mid++ {
		if (isScramble(s1[:mid], s2[:mid]) && isScramble(s1[mid:], s2[mid:])) || (isScramble(s1[:mid], s2[len(s2)-mid:]) && isScramble(s1[mid:], s2[:len(s2)-mid])) {
			m[Pair{s1, s2}] = true
			return true
		}
	}
	m[Pair{s1, s2}] = false
	return false
}

func isScramble(s1 string, s2 string) bool {
	return helper(map[Pair]bool{}, s1, s2)
}
func main() {
	fmt.Println(isScramble("xstjzkfpkggnhjzkpfjoguxvkbuopi", "xbouipkvxugojfpkzjhnggkpfkzjts"))
	fmt.Println(isScramble("great", "rgeat"))
}
