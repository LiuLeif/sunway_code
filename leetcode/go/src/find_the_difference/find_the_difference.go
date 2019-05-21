// 2017-12-29 14:47
package main

import "fmt"

func findTheDifference(s string, t string) byte {
	count := [26]int{}
	for i := 0; i < len(s); i++ {
		count[s[i]-'a']++
	}
	for i := 0; i < len(t); i++ {
		count[t[i]-'a']--
		if count[t[i]-'a'] < 0 {
			return t[i]
		}
	}
	return 0
}
func main() {
	fmt.Println(findTheDifference("abcd", "abcde"))
}
