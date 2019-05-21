// 2018-01-24 14:21
package main

import "fmt"

func getSumRecursive(a, b int) int {
	if b == 0 {
		return a
	}
	return getSumRecursive(a^b, (a&b)<<1)
}

func getSum(a int, b int) int {
	ret := 0
	carry := 0

	for i := uint(0); i < 64; i++ {
		a1, b1 := a&1, b&1
		a >>= 1
		b >>= 1
		ret = (a1^b1^carry)<<i | ret
		if a1&b1 == 1 || a1&carry == 1 || b1&carry == 1 {
			carry = 1
		} else {
			carry = 0
		}
	}
	return ret
}
func main() {
	fmt.Println(getSum(100, 100))
	fmt.Println(getSumRecursive(1, 2))
}
