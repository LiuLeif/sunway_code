// 2018-01-24 10:18
package main

import "fmt"

func isValid(x byte) bool {
	return (x >= 'A' && x <= 'Z') || (x >= 'a' && x <= 'z') || (x >= '0' && x <= '9')
}

func next(s string, i int) int {
	for i < len(s) && !isValid(s[i]) {
		i++
	}
	return i
}

func prev(s string, i int) int {
	for i >= 0 && !isValid(s[i]) {
		i--
	}
	return i
}

func lower(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		c = c - 'A' + 'a'
	}
	return c
}

func isPalindrome(s string) bool {
	lo := next(s, 0)
	hi := prev(s, len(s)-1)
	for lo < hi {
		if lower(s[lo]) != lower(s[hi]) {
			return false
		}
		lo = next(s, lo+1)
		hi = prev(s, hi-1)
	}
	return true
}
func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(isPalindrome("0P"))
}
