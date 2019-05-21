// 2018-10-31 11:09
package main

import (
	"fmt"
	"strings"
)

func isAdditiveNumber(num string) bool {
	add := func(a, b string) string {
		na := []rune(a)
		nb := []rune(b)
		for i := 0; i < len(na)/2; i++ {
			na[i], na[len(na)-i-1] = na[len(na)-i-1], na[i]
		}
		for i := 0; i < len(nb)/2; i++ {
			nb[i], nb[len(nb)-i-1] = nb[len(nb)-i-1], nb[i]
		}

		ret := []rune{}
		carry := 0
		for ia, ib := 0, 0; ia < len(na) || ib < len(nb); ia, ib = ia+1, ib+1 {
			va := 0
			if ia < len(na) {
				va = int(na[ia]) - '0'
			}
			vb := 0
			if ib < len(nb) {
				vb = int(nb[ib]) - '0'
			}
			tmp := va + vb + carry
			carry = tmp / 10
			tmp = tmp % 10
			ret = append(ret, rune(tmp+'0'))
		}
		if carry != 0 {
			ret = append(ret, rune(carry+'0'))
		}
		for i := 0; i < len(ret)/2; i++ {
			ret[i], ret[len(ret)-1-i] = ret[len(ret)-1-i], ret[i]
		}
		return string(ret)
	}

	var helper func(prev []string, num string) bool
	helper = func(prev []string, num string) bool {
		if len(num) == 0 {
			if len(prev) < 3 {
				return false
			}
			return true
		}

		if len(prev) < 2 {
			for i := 1; i < len(num); i++ {
				tmp := num[:i]
				if num[0] == '0' && i != 1 {
					break
				}
				prev = append(prev, tmp)
				if helper(prev, num[len(tmp):]) {
					return true
				}
				prev = prev[:len(prev)-1]
			}
		} else {
			tmp := add(prev[len(prev)-1], prev[len(prev)-2])
			if strings.HasPrefix(num, tmp) {
				prev = append(prev, tmp)
				if helper(prev, num[len(tmp):]) {
					return true
				}
				prev = prev[:len(prev)-1]
			}
		}
		return false
	}
	return helper([]string{}, num)
}

func main() {
	fmt.Println(isAdditiveNumber("11235"))
}
