// 2017-11-22 18:51
package main

import "fmt"

func trailingZeroes(n int) int {
	k := 5
	ret := 0
	for k <= n {
		ret += n / k
		k = k * 5
	}
	return ret
}
func main() {
	fmt.Println(trailingZeroes(200))
}
