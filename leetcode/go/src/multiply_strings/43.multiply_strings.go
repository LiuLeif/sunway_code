// 2018-01-30 13:18
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func multiply(num1 string, num2 string) string {
	m, n := len(num1), len(num2)
	ret := make([]int, m+n)
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			a := int(num1[i] - '0')
			b := int(num2[j] - '0')
			mul := a*b + ret[i+j+1]
			ret[i+j+1] = mul % 10
			ret[i+j] += mul / 10
		}
	}

	s := ""
	for i := 0; i < m+n; i++ {
		s += strconv.Itoa(ret[i])
	}
	s = strings.TrimLeft(s, "0")
	if len(s) == 0 {
		return "0"
	} else {
		return s
	}
}
func main() {
	fmt.Println(multiply("123", "4"))
}
