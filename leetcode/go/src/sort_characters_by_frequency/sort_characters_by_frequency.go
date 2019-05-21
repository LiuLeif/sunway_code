// 2017-12-11 16:36
package main

import "fmt"
import "strings"

func frequencySort(s string) string {
	m := map[byte]int{}
	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}
	buckets := make([]string, len(s)+1)
	for k, v := range m {
		buckets[v] += strings.Repeat(string(k), v)
	}
	ret := ""
	for i := len(s); i > 0; i-- {
		ret += buckets[i]
	}

	return ret
}
func main() {
	fmt.Println(frequencySort("aabacd"))
}
