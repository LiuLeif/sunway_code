// 2018-02-24 13:49
package main

import "fmt"

func isCapital(c byte) bool {
	return c >= 'A' && c <= 'Z'
}

func detectCapitalUse(word string) bool {
	if len(word) == 0 {
		return true
	}
	isUpper := false
	isLower := false
	for i := 1; i < len(word); i++ {
		if isCapital(word[i]) {
			if isLower {
				return false
			}
			isUpper = true
		} else {
			if isUpper {
				return false
			}
			isLower = true
		}
	}
	if !isCapital(word[0]) && isUpper {
		return false
	}
	return true
}
func main() {
	fmt.Println(detectCapitalUse(""))
}
