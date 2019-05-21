// 2018-01-02 17:35
package main

import (
	"fmt"
	"math"
)

func reachNumber(n int) int {
	n = int(math.Abs(float64(n)))
	ret, sum := 1, 0
	for sum < n || (sum-n)%2 == 1 {
		sum += ret
		ret++
	}
	return ret - 1
}
func main() {
	// 1: -1, 1
	// 2: -3, -1, 1, 3
	// 3: -6, -4, -2, 0, 2, 4, 6
	// 4: -10,-8,-6, -4, -2, 0, 2, 4, 6, 8, 10
	// 5: -15,-13, -11, -9,-7,-5,-3,1,3,5,7,9,11,13,15
	// 6: -21,-19,-17,-15,-13... 21
	// 7: -28
	fmt.Println(reachNumber(5))
}
