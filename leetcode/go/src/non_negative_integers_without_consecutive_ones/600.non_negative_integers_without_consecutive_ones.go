// 2018-12-06 15:26
package main

import "fmt"

func findIntegers(num int) int {
	// 0 0
	// 1 1
	// 2 10
	// 3 11    <--
	// 4 100
	// 5 101 100 101
	// 6 110   <--
	// 7 111   <--
	// 111 = 100 11
	// 0~11111 = 0~1111+10000~11111
	//                  10000~10111
	x, y := 1, 2
	fib := []int{x, y}
	for i := 0; i < 32; i++ {
		x, y = y, x+y
		fib = append(fib, y)
	}
	ret := 0
	prevIsOne := false
	for k := 30; k >= 0; k-- {
		if num&(1<<uint32(k)) != 0 {
			ret += fib[k]
			if prevIsOne {
				return ret
			}
			prevIsOne = true
		} else {
			prevIsOne = false
		}
	}
	return ret + 1
}

func main() {
	fmt.Println(findIntegers(3))
}
