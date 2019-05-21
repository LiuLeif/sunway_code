// 2018-01-10 16:33
package main

import (
	"fmt"
	"math"
)

func isPowerOfThree(n int) bool {
	return n > 0 && int(math.Pow(3, 19))%n == 0
}
func main() {
	fmt.Println(isPowerOfThree(21))
}
