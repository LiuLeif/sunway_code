// 2018-01-18 10:58
package main

import (
	"fmt"
	"strings"
)

func wordPattern(pattern string, str string) bool {
	if len(pattern) == 0 && len(str) == 0 {
		return true
	}
	ma := map[byte]string{}
	mb := map[string]byte{}

	strs := strings.Split(str, " ")
	if len(pattern) != len(strs) {
		return false
	}
	for i := 0; i < len(pattern); i++ {
		a, b := pattern[i], strs[i]
		if ma[a] != "" && ma[a] != b {
			return false
		}
		if mb[b] != 0 && mb[b] != a {
			return false
		}
		if ma[a] == "" {
			ma[a] = b
		}
		if mb[b] == 0 {
			mb[b] = a
		}
	}
	return true
}
func main() {
	fmt.Println(wordPattern("", ""))
}
