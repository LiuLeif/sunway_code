// 2018-02-06 13:18
package main

import (
	"fmt"
	"strings"
)

func intToRoman(num int) string {
	// I（1）   X（10）   C（100）   M（1000）
	//     V（5）    L（50）   D（500）
	// 3999
	// 1980 MCMLXXX
	ret := []string{}
	romans := "IVXLCDMZZ"
	base := 0
	for num > 0 {
		n := num % 10
		num /= 10

		if n >= 1 && n <= 3 {
			ret = append(ret, strings.Repeat(string(romans[base]), n))
		} else if n == 4 {
			ret = append(ret, string(romans[base])+string(romans[base+1]))
		} else if n == 5 {
			ret = append(ret, string(romans[base+1]))
		} else if n >= 6 && n <= 8 {
			ret = append(ret, string(romans[base+1])+strings.Repeat(string(romans[base]), n-5))
		} else if n == 9 {
			ret = append(ret, string(romans[base])+string(romans[base+2]))
		}
		base += 2
	}
	for i := 0; i < len(ret)/2; i++ {
		ret[i], ret[len(ret)-1-i] = ret[len(ret)-1-i], ret[i]
	}
	return strings.Join(ret, "")
}
func main() {
	fmt.Println(intToRoman(3999))
}
