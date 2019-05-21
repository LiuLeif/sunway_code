// 2018-03-02 10:10
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func rotatedDigits(N int) int {
	same := []int{1, 2, 2, 2, 2, 2, 2, 2, 3, 3}
	diff := []int{0, 0, 1, 1, 1, 2, 3, 3, 3, 4}
	var calc func(s string) (int, int)
	calc = func(s string) (int, int) {
		if len(s) == 1 {
			return same[s[0]-'0'], diff[s[0]-'0']
		}
		s9, d9 := calc(strings.Repeat("9", len(s)-1))
		sN, dN := calc(s[1:])
		if s[0] == '0' {
			return sN, dN
		}

		lead := s[0] - '0'
		sRet := same[lead-1] * s9
		switch lead {
		case 0, 1, 8:
			sRet += sN
		}
		dRet := (same[lead-1]+diff[lead-1])*d9 + diff[lead-1]*s9
		switch lead {
		case 0, 1, 8:
			dRet += dN
		case 2, 5, 6, 9:
			dRet += sN + dN
		}
		return sRet, dRet
	}
	_, ret := calc(strconv.Itoa(N))
	return ret
}
func main() {
	fmt.Println(rotatedDigits(502))
}
