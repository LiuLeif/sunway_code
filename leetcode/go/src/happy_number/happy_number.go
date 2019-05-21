// 2017-11-17 11:20
package main

import "fmt"

func nextNumber(n int) int {
	ret := 0
	for n != 0 {
		t := n % 10
		n = n / 10
		ret += t * t
	}
	return ret
}

func isHappy(n int) bool {
	x, y := n, n
	for {
		x = nextNumber(x)
		y = nextNumber(y)
		y = nextNumber(y)
		if x == 1 {
			return true
		}
		if x == y {
			return false
		}
	}

}
func main() {
	fmt.Println(isHappy(68))
}
