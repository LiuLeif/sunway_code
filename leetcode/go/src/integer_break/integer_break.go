// 2018-01-09 11:41
package main

import "fmt"

func integerBreak(n int) int {
	switch n {
	case 1:
		return 1
	case 2:
		return 1
	case 3:
		return 2
	}
	ret := 1
	for n > 4 {
		ret *= 3
		n -= 3
	}
	ret *= n
	return ret
}
func main() {
	fmt.Println(integerBreak(10))
}
