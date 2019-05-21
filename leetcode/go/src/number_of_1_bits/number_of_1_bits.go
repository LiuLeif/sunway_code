// 2017-11-20 16:14
package main

import "fmt"

func hammingWeight(num int) int {
	ret := 0
	for num != 0 {
		num &= (num - 1)
		ret++
	}
	return ret
}

func main() {
	fmt.Println(hammingWeight(11))
}
