// 2017-11-20 16:14
package main

import "fmt"

func hammingDistance(x int, y int) int {
	ret := 0
	for i := uint(0); i < 32; i++ {
		if (x & (1 << i)) != (y & (1 << i)) {
			ret++
		}
	}
	return ret
}
func main() {
	fmt.Println(hammingDistance(1, 1))
}
