// 2018-11-20 10:05
package main

import (
	"fmt"
	"strconv"
)

func countDigitOne(n int) int {
	if n < 0 {
		return 0
	}
	S := strconv.Itoa(n)
	ret := 0
	for i := 0; i < len(S); i++ {
		c := S[i]
		left := S[:i]
		right := S[i+1:]
		if c > '1' {
			// change right to 9
			newRight := "``"
			for i := 0; i < len(right); i++ {
				newRight += string('9')
			}
			tmp, _ := strconv.Atoi(left + newRight)
			ret += tmp + 1
		} else if c == '1' {
			tmp, _ := strconv.Atoi(left + right)
			ret += tmp + 1
		} else {
			tmp, _ := strconv.Atoi(left + right)
			ret += tmp + 1
			tmp, _ = strconv.Atoi(right)
			ret -= tmp + 1
		}
	}
	return ret
}

func main() {
	fmt.Println(countDigitOne(2)) // 12

}
