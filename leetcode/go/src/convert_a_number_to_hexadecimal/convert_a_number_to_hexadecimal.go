// 2018-01-22 10:35
package main

import (
	"fmt"
)

func toHex2(num int) string {
	ret := []byte{}
	m := "0123456789abcdef"
	unum := uint32(num)
	for unum != 0 {
		t := unum & 0xf
		ret = append(ret, m[t])
		unum = unum >> 4
	}
	for i := 0; i < len(ret)/2; i++ {
		ret[i], ret[len(ret)-1-i] = ret[len(ret)-1-i], ret[i]
	}

	return string(ret)
}

func toHex(num int) string {
	ret := []byte{}
	unum := uint32(num)
	for unum >= 16 {
		t := unum % 16
		switch {
		case t < 10:
			ret = append(ret, byte(t+'0'))
		default:
			ret = append(ret, byte(t-10+'a'))
		}
		unum /= 16
	}
	switch {
	case unum < 10:
		ret = append(ret, byte(unum+'0'))
	default:
		ret = append(ret, byte(unum-10+'a'))
	}

	for i := 0; i < len(ret)/2; i++ {
		ret[i], ret[len(ret)-1-i] = ret[len(ret)-1-i], ret[i]
	}
	return string(ret)
}
func main() {
	fmt.Println(toHex2(26))
}
