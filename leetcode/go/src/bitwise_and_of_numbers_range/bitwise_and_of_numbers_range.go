// 2017-11-14 18:54
package main

import "fmt"

func rangeBitwiseAnd(m int, n int) int {
	for m < n {
		n = n & (n - 1)
	}
	return n
}
func main() {
	fmt.Println(rangeBitwiseAnd(5, 7))
}
