// 2017-12-27 17:05
package main

import "fmt"

func checkInclusion(s1 string, s2 string) bool {
	l1, l2 := len(s1), len(s2)
	if l2 < l1 {
		return false
	}
	count1 := [26]int{}
	for i := 0; i < len(s1); i++ {
		count1[s1[i]-'a']++
	}
	count2 := [26]int{}
	for i := 0; i < len(s1); i++ {
		count2[s2[i]-'a']++
	}
	if count1 == count2 {
		return true
	}
	for i := 1; i < l2-l1+1; i++ {
		count2[s2[i-1]-'a']--
		count2[s2[i+l1-1]-'a']++
		if count1 == count2 {
			return true
		}
	}
	return false
}
func main() {
	fmt.Println(checkInclusion("ab", "ba"))
}
