// 2018-02-13 14:29
package main

import (
	"fmt"
	"strings"
)

func reverseString(s string) string {
	b := []byte(s)
	for i := 0; i < len(s)/2; i++ {
		b[i], b[len(b)-1-i] = b[len(b)-1-i], b[i]
	}
	return string(b)
}

func reverseWords(s string) string {
	strs := strings.Split(s, " ")
	for i := 0; i < len(strs); i++ {
		strs[i] = reverseString(strs[i])
	}

	return strings.Join(strs, " ")
}
func main() {
	fmt.Println(reverseWords("ax bx"))
}
