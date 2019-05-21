// 2017-12-13 13:52
package main

import (
	"fmt"
)

func isPowerOfTwo(n int) bool {
	if n <= 0 {
		return false
	}
	x := 1073741824
	return x%n == 0
}
func main() {
	fmt.Println(isPowerOfTwo(1))
}
