// 2017-12-15 17:52
package main

import "fmt"

func addDigits(num int) int {
	// 0 1 2 3 4 5 6 7 8 9 10 11 12 ...      18 19 20
	// 0 1 2 3 4 5 6 7 8 9 1  2  3  ...      9  1  2
	return 1 + (num-1)%9
}
func main() {
	fmt.Println(addDigits(38))
}
