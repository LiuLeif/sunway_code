// 2018-02-06 11:50
package main

import (
	"fmt"
	"math"
)

func reverse(s string) string {
	b := []byte(s)
	for i := 0; i < len(s)/2; i++ {
		b[i], b[len(b)-1-i] = b[len(b)-1-i], b[i]
	}
	return string(b)
}

func max(nums ...int) int {
	ret := math.MinInt32
	for _, n := range nums {
		if ret < n {
			ret = n
		}
	}
	return ret
}

func addBinary(a string, b string) string {
	ret := ""
	a = reverse(a)
	b = reverse(b)
	carry := 0
	for i := 0; i < max(len(a), len(b)); i++ {
		x := 0
		y := 0
		if i < len(a) {
			x = int(a[i] - '0')
		}
		if i < len(b) {
			y = int(b[i] - '0')
		}
		sum := x + y + carry
		ret += string(sum%2 + '0')
		carry = sum / 2
	}
	if carry != 0 {
		ret += string(carry + '0')
	}
	return reverse(ret)
}
func main() {
	fmt.Println(addBinary("0", "1"))
}
