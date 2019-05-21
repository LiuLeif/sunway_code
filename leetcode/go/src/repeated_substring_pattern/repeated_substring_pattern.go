// 2018-02-01 10:34
package main

import "fmt"

func repeatedSubstringPattern(s string) bool {
NextDividor:
	for dividor := 2; dividor <= len(s); dividor++ {
		if len(s)%dividor != 0 {
			continue
		}
		step := len(s) / dividor
		subs := s[:step]
		for i := 1; i < dividor; i++ {
			if s[step*i:step*i+step] != subs {
				continue NextDividor
			}
		}
		return true
	}
	return false
}
func main() {
	fmt.Println(repeatedSubstringPattern("abcabcabcabc"))
	fmt.Println(repeatedSubstringPattern("aaa"))
}
