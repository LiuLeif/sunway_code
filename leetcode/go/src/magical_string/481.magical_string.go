// 2018-11-02 09:54
package main

import "fmt"

func magicalString(n int) int {
	// 122
	// 122
	s := "122"
	for i := 2; i < n; i++ {
		if s[i] == '2' {
			// 2
			if s[len(s)-1] == '1' {
				s += "22"
			} else {
				s += "11"
			}
		} else {
			// 1
			if s[len(s)-1] == '1' {
				s += "2"
			} else {
				s += "1"
			}
		}
	}
	ret := 0
	for i := 0; i < n; i++ {
		if s[i] == '1' {
			ret++
		}
	}
	return ret
}

func main() {
	fmt.Println(magicalString(6))
}
