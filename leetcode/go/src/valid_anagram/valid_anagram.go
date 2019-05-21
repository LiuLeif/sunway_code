// 2017-12-22 15:50
package main

import "fmt"

func isAnagram(s string, t string) bool {
	bucket := [26]int{}
	for i := 0; i < len(s); i++ {
		bucket[s[i]-'a']++
	}
	for i := 0; i < len(t); i++ {
		bucket[t[i]-'a']--
	}
	for i := 0; i < 26; i++ {
		if bucket[i] != 0 {
			return false
		}
	}
	return true
}
func main() {
	fmt.Println(isAnagram("a", "ab"))
}
