// 2017-11-23 14:32
package main

import "fmt"

func strStr(haystack string, needle string) int {
	if len(needle) > len(haystack) {
		return -1
	}

	for i := 0; i < len(haystack)-len(needle)+1; i++ {
		if haystack[i:len(needle)+i] == needle {
			return i
		}
	}

	return -1
}
func main() {
	fmt.Println(strStr("missi", "issib"))
}
