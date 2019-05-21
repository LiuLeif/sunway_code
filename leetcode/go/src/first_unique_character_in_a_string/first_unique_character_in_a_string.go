// 2018-01-24 13:17
package main

import "fmt"

func firstUniqChar(s string) int {
	count := make([]int, len(s))
	m := map[byte]int{}
	for i := 0; i < len(s); i++ {
		if _, ok := m[s[i]]; !ok {
			m[s[i]] = i
		}
		count[m[s[i]]]++
	}
	for i := 0; i < len(count); i++ {
		if count[i] == 1 {
			return i
		}
	}
	return -1
}
func main() {
	fmt.Println(firstUniqChar("aab"))
}
