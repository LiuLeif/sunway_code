// 2017-12-15 17:52
package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	y := x
	if y < 0 {
		y = -x
	}
	ret := 0
	for y != 0 {
		ret = ret*10 + y%10
		y = y / 10
	}
	if x < 0 {
		ret = -ret
	}
	if ret > math.MaxInt32 || ret < math.MinInt32 {
		return 0
	}
	return ret
}
func main() {
	fmt.Println(reverse(123))
}
