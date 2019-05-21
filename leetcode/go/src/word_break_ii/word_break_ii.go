// 2018-01-08 14:22
package main

import (
	"fmt"
	"strings"
)

func helper(s string, dict []string, m map[string][]string) []string {
	ret := []string{}
	if len(s) == 0 {
		return []string{""}
	}
	if m[s] != nil {
		return m[s]
	}
	for i := 0; i < len(dict); i++ {
		if strings.HasPrefix(s, dict[i]) {
			x := helper(s[len(dict[i]):], dict, m)
			for j := 0; j < len(x); j++ {
				if len(x[j]) == 0 {
					ret = append(ret, dict[i])
				} else {
					ret = append(ret, dict[i]+" "+x[j])
				}
			}
		}
	}
	m[s] = ret
	return ret
}

func wordBreak(s string, dict []string) []string {
	return helper(s, dict, map[string][]string{})
}
func main() {
	dict := []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"}
	fmt.Println(wordBreak("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaabaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", dict))
	dict = []string{"cat", "cats", "and", "sand", "dog"}
	fmt.Println(wordBreak("catsanddog", dict))
}
