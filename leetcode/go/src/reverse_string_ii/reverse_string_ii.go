// 2018-02-24 10:03
package main

import (
	"fmt"
	"math"
)

func min(nums ...int) int {
	ret := math.MaxInt32
	for _, n := range nums {
		if ret > n {
			ret = n
		}
	}
	return ret
}

func reverse(buff []byte) {
	for i := 0; i < len(buff)/2; i++ {
		buff[i], buff[len(buff)-i-1] = buff[len(buff)-i-1], buff[i]
	}
}

func reverseStr(s string, k int) string {
	buff := []byte(s)
	for i := 0; i < len(s); i += 2 * k {
		reverse(buff[i:min(i+k, len(s))])
	}
	return string(buff)
}

func main() {
	fmt.Println(reverseStr("abcdefg", 2))
}
