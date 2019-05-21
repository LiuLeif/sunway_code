// 2017-11-15 16:16
package main

import "fmt"

func mySqrt(x int) int {
	low, high := 0, x
	for low <= high {
		mid := low + (high-low)/2
		tmp := mid * mid
		if tmp == x {
			return mid
		}
		if tmp > x {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return high
}

func main() {
	fmt.Println(mySqrt(9))
}
