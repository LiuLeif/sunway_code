// 2018-01-19 11:13
package main

import (
	"fmt"
)

func convertToBase7(num int) string {
	negative := false
	if num < 0 {
		num = -num
		negative = true
	}
	ret := []byte{}
	for num >= 7 {
		ret = append(ret, byte(num%7+'0'))
		num /= 7
	}
	ret = append(ret, byte(num+'0'))
	for i := 0; i < len(ret)/2; i++ {
		ret[i], ret[len(ret)-1-i] = ret[len(ret)-1-i], ret[i]
	}
	if negative {
		return "-" + string(ret)
	} else {
		return string(ret)
	}
}
func main() {
	fmt.Println(convertToBase7(12))
}
