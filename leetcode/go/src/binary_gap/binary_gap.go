// 2018-10-16 09:59
package main

import "fmt"

func binaryGap(N int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	prev := -1
	index := 0
	ret := 0
	for N != 0 {
		if N&1 == 1 {
			if prev != -1 {
				ret = max(ret, index-prev)
			}
			prev = index
		}
		index++
		N >>= 1
	}
	return ret
}
func main() {
	fmt.Println(binaryGap(8))
}
