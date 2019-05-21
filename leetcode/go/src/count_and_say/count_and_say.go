// 2018-01-23 17:18
package main

import (
	"fmt"
	"strconv"
)

func countAndSay(n int) string {
	s := "1"
	for i := 1; i < n; i++ {
		// 1211
		count := 1
		t := ""
		for j := 1; j < len(s); j++ {
			if s[j] == s[j-1] {
				count++
			} else {
				t += strconv.Itoa(count) + string(s[j-1])
				count = 1
			}
		}
		t += strconv.Itoa(count) + string(s[len(s)-1])
		s = t
	}
	return s
}
func main() {
	fmt.Println(countAndSay(4))
}
