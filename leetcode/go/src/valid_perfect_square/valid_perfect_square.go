// 2017-11-14 18:54
package main

import "fmt"

func isPerfectSquareBS(num int) bool {
	low, high := 0, num
	for low <= high {
		mid := low + (high-low)/2
		x := mid * mid
		if x == num {
			return true
		}
		if x < num {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return false
}

func isPerfectSquare(num int) bool {
	// 1,3,5,7,9
	x := 1
	total := 0
	for total < num {
		total += x
		x += 2
	}
	return total == num
}

func main() {
	fmt.Println(isPerfectSquareBS(4))
}
