// 2018-02-09 15:32
package main

import "fmt"

func integerReplacement(n int) int {
	// 1011 1100 1010
	ret := 0
	for n != 1 {
		ret++
		if n&1 == 0 {
			n >>= 1
		} else {
			if n&3 == 3 && n != 3 {
				n++
			} else {
				n--
			}
		}
	}
	return ret
}

func main() {
	fmt.Println(integerReplacement(3))
}
