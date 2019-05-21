// 2017-12-07 14:07
package main

import (
	"fmt"
	"strings"
)

func reverse(s string) string {
	b := []byte(s)
	for i := 0; i < len(s)/2; i++ {
		b[i], b[len(s)-i-1] = b[len(s)-i-1], b[i]
	}
	return string(b)
}

func shortestPalindrome(s string) string {
	if len(s) == 0 {
		return s
	}
	//  xyxy
	// yxyx
	r := reverse(s)
	for i := 0; i < len(r); i++ {
		if strings.HasPrefix(s, r[i:]) {
			return r[:i] + s
		}
	}
	return s
}
func main() {
	fmt.Println(shortestPalindrome("abab"))
}
