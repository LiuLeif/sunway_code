// 2018-02-12 10:40
package main

import (
	"fmt"
	"strings"
)

type pair struct {
	ret   string
	count int
}

func decodeString(s string) string {
	ret := ""
	count := 0
	stack := []pair{}
	for i := 0; i < len(s); i++ {
		c := s[i]
		switch {
		case c == '[':
			stack = append(stack, pair{ret, count})
			ret = ""
			count = 0
		case c == ']':
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ret = top.ret + strings.Repeat(ret, top.count)
		case c >= '0' && c <= '9':
			count = count*10 + int(c-'0')
		default:
			ret += string(c)
		}
	}
	return ret
}
func main() {
	fmt.Println(decodeString("10[bc]"))
}
