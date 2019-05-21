// 2017-11-20 16:14
package main

import "fmt"

func hasAlternatingBits(n int) bool {
	x := n
	y := n >> 1
	if x&y != 0 {
		return false
	}
	z := x | y
	return z&(z+1) == 0

}

func main() {
	fmt.Println(hasAlternatingBits(5))

}
