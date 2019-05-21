// 2017-12-20 14:33
package main

import "fmt"

func isSubsequence(s string, t string) bool {
	si, ti := 0, 0

	for si < len(s) && ti < len(t) {
		if t[ti] == s[si] {
			si += 1
			ti += 1
		} else {
			ti += 1
		}
	}

	return si == len(s)
}
func main() {
	fmt.Println(isSubsequence("abc", "deaxbxac"))
}
