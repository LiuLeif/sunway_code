// 2017-12-26 16:09
package main

import "fmt"

func findComplement(num int) int {
	n := 0
	t := num
	for t != 0 {
		t = t >> 1
		n++
	}
	mask := (1 << uint(n)) - 1
	return (^num) & mask
}
func main() {
	fmt.Println(findComplement(2))
}
