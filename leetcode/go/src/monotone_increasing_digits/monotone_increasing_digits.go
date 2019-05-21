// 2018-02-27 10:52
package main

import (
	"fmt"
	"strconv"
)

func monotoneIncreasingDigits(N int) int {
	s := []byte(strconv.Itoa(N))
	s = append(s, '9')
	left := 0
	for i := 1; i < len(s); i++ {
		if s[i] > s[left] {
			left = i
		} else if s[i] < s[left] {
			break
		}
	}
	s[left]--
	for i := left + 1; i < len(s); i++ {
		s[i] = '9'
	}
	s = s[:len(s)-1]
	ret, _ := strconv.Atoi(string(s))
	return ret
}
func main() {
	// 423
	// 142
	fmt.Println(monotoneIncreasingDigits(1322))
	fmt.Println(monotoneIncreasingDigits(0))
}
