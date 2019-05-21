// 2018-11-12 10:44
package main

import "fmt"

func check(S, s string) bool {
	// heeellooo
	// hello
	j := 0
	for i := 0; i < len(S); i++ {
		if j < len(s) && S[i] == s[j] {
			j++
		} else if i > 0 && i < len(S)-1 && S[i-1] == S[i] && S[i] == S[i+1] {
		} else if i > 1 && S[i] == S[i-1] && S[i] == S[i-2] {
		} else {
			return false
		}
	}
	return j == len(s)
}

func expressiveWords(S string, words []string) int {
	ret := 0
	for _, s := range words {
		if check(S, s) {
			ret++
		}
	}
	return ret
}

func main() {
	fmt.Println(expressiveWords("heeellooo", []string{"hello", "hi", "helo"}))
}
