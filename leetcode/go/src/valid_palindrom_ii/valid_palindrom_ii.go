// 2017-11-14 18:54
package main

import "fmt"

func isPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}

func validPalindrome(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] != s[len(s)-i-1] {
			return isPalindrome(s[i+1:len(s)-i]) || isPalindrome(s[i:len(s)-i-1])
		}
	}
	return true
}
func main() {
	fmt.Println(validPalindrome("aaax"))
}
