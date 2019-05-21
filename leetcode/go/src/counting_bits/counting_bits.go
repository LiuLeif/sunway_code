// 2017-11-20 16:14
package main

import "fmt"

func countBits(num int) []int {
	ret := make([]int, num+1)
	ret[0] = 0
	for i := 1; i < num+1; i++ {
		if (i & 1) == 0 {
			ret[i] = ret[i>>1]
		} else {
			ret[i] = ret[i>>1] + 1
		}
	}
	return ret
}
func main() {
	fmt.Println(countBits(19))
}
