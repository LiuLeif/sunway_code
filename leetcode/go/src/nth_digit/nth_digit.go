// 2018-01-04 17:04
package main

import "fmt"

func findNthDigit(n int) int {
	k := 1
	digit := 9
	all := 9
	for all < n {
		k += 1
		digit = digit * 10
		all = all + digit*k
	}
	remain := n - all + digit*k
	// get remain from kth level
	x := (remain - 1) / k
	y := (remain - 1) % k
	firstNumber := 1
	if k == 1 {
		firstNumber = 1
	} else {
		for i := 0; i < k-1; i++ {
			firstNumber *= 10
		}
	}
	number := firstNumber + x
	for i := 0; i < k-y-1; i++ {
		number /= 10
	}
	return number % 10
}
func main() {
	// 1234567891011121314...
	// 0~9 = 9
	// 10~99 = 90
	// 100~999 = 900
	fmt.Println(findNthDigit(110))
}
