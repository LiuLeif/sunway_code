// 2018-01-22 14:23
package main

import (
	"fmt"
	"math"
)

func bulbSwitch(n int) int {
	return int(math.Sqrt(float64(n)))
}
func main() {
	// n=4
	// 1: 1 1 1 1
	// 2: 1 0 1 0
	// 3: 1 0 0 0
	// 4: 1 0 0 1
	// n=5
	// 1: 1 1 1 1 1
	// 2: 1 0 1 0 1
	// 3: 1 0 0 0 1
	// 4: 1 0 0 1 1
	// 5: 1 0 0 1 0
	// 所有 squre number 都是 on 的, 例如 1,4,9,16... 因为它们的 divider 是奇数个
	fmt.Println(bulbSwitch(10))
}
