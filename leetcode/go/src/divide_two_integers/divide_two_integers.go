// 2018-01-26 11:41
package main

import (
	"fmt"
	"math"
)

func divide(dividend int, divisor int) int {
	if divisor == 0 {
		return math.MaxInt32
	}
	negative := false
	if dividend < 0 {
		dividend = -dividend
		negative = !negative
	}
	if divisor < 0 {
		divisor = -divisor
		negative = !negative
	}

	ret := 0
	for dividend >= divisor {
		x := divisor
		n := 1
		for dividend >= x {
			x <<= 1
			n <<= 1
		}
		ret += n >> 1
		dividend -= x >> 1
	}
	if negative {
		ret = -ret
	}
	if ret > math.MaxInt32 {
		return math.MaxInt32
	}
	return ret
}
func main() {
	fmt.Println(divide(11, -10))
}
