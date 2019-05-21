// 2018-02-06 13:51
package main

import "fmt"

func romanToInt(s string) int {
	// I（1）   X（10）   C（100）   M（1000）
	//     V（5）    L（50）   D（500）
	m := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	ret := 0
	max := 0
	// 1000 100 100 1000 50 10 10 10
	for i := len(s) - 1; i >= 0; i-- {
		if m[s[i]] >= max {
			max = m[s[i]]
			ret += m[s[i]]
		} else {
			ret -= m[s[i]]
		}
	}
	return ret
}
func main() {
	fmt.Println(romanToInt("X"))
}
