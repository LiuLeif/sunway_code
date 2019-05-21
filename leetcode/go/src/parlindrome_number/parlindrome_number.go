// 2017-11-14 18:54
package main

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	y := x
	k := 1
	for y != 0 {
		y /= 10
		k *= 10
	}
	k /= 10

	for k > 1 {
		left := x / k
		x %= k
		k /= 10
		right := x % 10
		if left != right {
			return false
		}
		k /= 10
		x /= 10
	}

	return true
}
func main() {
	fmt.Println(isPalindrome(123221))
}
