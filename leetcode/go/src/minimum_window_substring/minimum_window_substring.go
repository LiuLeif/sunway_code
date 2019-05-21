// 2018-02-05 13:52
package main

import "fmt"

func minWindow(s string, t string) string {
	if len(s) == 0 || len(t) == 0 {
		return ""
	}
	needed := map[byte]int{}
	have := map[byte]int{}
	for i := 0; i < len(t); i++ {
		needed[t[i]]++
	}
	lo, hi := 0, -1
	seek := func(c byte) {
		if have[c] >= needed[c] {
			return
		}
		for hi = hi + 1; hi < len(s); hi++ {
			if needed[s[hi]] != 0 {
				have[s[hi]]++
				if s[hi] == c {
					break
				}
			}
		}
	}

	for i := 0; i < len(t); i++ {
		seek(t[i])
	}

	if hi >= len(s) {
		return ""
	}

	ret := s[lo : hi+1]
	for ; hi < len(s); lo++ {
		ret = min(ret, s[lo:hi+1])
		if needed[s[lo]] == 0 {
			continue
		}
		have[s[lo]]--
		seek(s[lo])
	}
	return ret
}

func min(s1, s2 string) string {
	if len(s1) <= len(s2) {
		return s1
	}
	return s2
}

func main() {
	fmt.Println(minWindow("ADOBECODEBANC", "AEC"))
}
