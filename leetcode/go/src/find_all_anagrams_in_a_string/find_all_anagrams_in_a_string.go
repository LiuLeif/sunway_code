// 2017-11-14 18:54
package main

import (
	"fmt"
)

func findAnagrams(s string, p string) []int {
	ret := []int{}
	if len(s) < len(p) {
		return ret
	}
	count := [26]int{}
	for i := 0; i < len(p); i++ {
		count[p[i]-'a']++
	}
	count2 := [26]int{}
	for i := 0; i < len(p); i++ {
		count2[s[i]-'a']++
	}
	if count2 == count {
		ret = append(ret, 0)
	}
	for i := 1; i < len(s)-len(p)+1; i++ {
		count2[s[i-1]-'a']--
		count2[s[i+len(p)-1]-'a']++
		if count2 == count {
			ret = append(ret, i)
		}
	}
	return ret
}
func main() {
	fmt.Println(findAnagrams("abxab", "ba"))
}
