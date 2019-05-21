// 2017-11-14 18:54
package main

import "fmt"

func grayCode(n int) []int {
	ret := []int{0, 1}

	for i := uint32(1); i < uint32(n); i++ {
		size := len(ret)
		for j := size - 1; j >= 0; j-- {
			ret = append(ret, ret[j]|(1<<i))
		}
	}

	return ret
}
func main() {
	fmt.Println(grayCode(2))
}
