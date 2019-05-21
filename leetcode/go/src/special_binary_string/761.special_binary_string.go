// 2018-11-09 15:42
package main

import (
	"fmt"
	"sort"
	"strings"
)

func makeLargestSpecial(S string) string {
	// 11011000
	// 1 10 1100 0
	// (()(()))
	if len(S) == 0 {
		return ""
	}
	count, start := 0, 0
	v := []string{}
	for i := 0; i < len(S); i++ {
		if S[i] == '1' {
			count++
		} else {
			count--
		}
		if count == 0 {
			if start+1 == i {
				v = append(v, "10")
			} else {
				v = append(v, "1"+makeLargestSpecial(S[start+1:i])+"0")
			}
			start = i + 1
		}
	}
	sort.Sort(sort.Reverse(sort.StringSlice(v)))
	return strings.Join(v, "")
}

func main() {
	fmt.Println(makeLargestSpecial("11011000"))
}
