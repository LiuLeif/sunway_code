// 2018-01-26 16:20
package main

import "fmt"

func isOk(s string, step int, counter map[string]int) bool {
	curr := map[string]int{}
	for i := 0; i < len(s)/step; i++ {
		x := s[i*step : i*step+step]
		curr[x]++
	}
	for k, v := range curr {
		if counter[k] != v {
			return false
		}
	}

	return true
}

func findSubstring(s string, words []string) []int {
	ret := []int{}
	counter := map[string]int{}
	for _, word := range words {
		counter[word]++
	}

	totalLen := len(words) * len(words[0])
	for i := 0; i < len(s)-totalLen+1; i++ {
		if isOk(s[i:i+totalLen], len(words[0]), counter) {
			ret = append(ret, i)
		}
	}

	return ret
}
func main() {
	fmt.Println(findSubstring("barfoothefoobarman", []string{"foo"}))
}
