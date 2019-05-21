// 2017-12-13 13:52
package main

import "fmt"

func isPowerOfFour(num int) bool {
	return num > 0 && num&(num-1) == 0 && (num-1)%3 == 0
}
func main() {
	fmt.Println(isPowerOfFour(16))
	fmt.Println(isPowerOfFour(8))
	fmt.Println(isPowerOfFour(0))
}
