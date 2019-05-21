// 2018-02-01 16:20
package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	ret := 0
	for x > 0 {
		ret = ret*10 + x%10
		x /= 10
	}
	return ret
}

func largestPalindrome(n int) int {
	// 99*99
	if n == 1 {
		return 9
	}
	maxValue := int(math.Pow10(n) - 1)
	for left := maxValue; left > maxValue/10; left-- {
		parlindrom := left*int(math.Pow10(n)) + reverse(left)
		for t := maxValue; t*t > parlindrom; t-- {
			if parlindrom%t == 0 {
				return parlindrom % 1337
			}
		}

	}
	return 0
}
func main() {
	fmt.Println(largestPalindrome(7))
}
