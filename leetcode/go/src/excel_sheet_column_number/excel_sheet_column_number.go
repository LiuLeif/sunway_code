// 2018-02-07 14:42
package main

import "fmt"

func titleToNumber(s string) int {
	// A:1 B:2 ... Z:26 AA:27 AB:28
	ret := 0
	x := 1
	for i := len(s) - 1; i >= 0; i-- {
		ret += int(s[i]-'A'+1) * x
		x *= 26
	}
	return ret
}
func main() {
	fmt.Println(titleToNumber("B"))
}
