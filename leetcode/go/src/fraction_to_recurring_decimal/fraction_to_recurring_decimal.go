// 2018-02-08 16:06
package main

import (
	"fmt"
	"strconv"
)

func fractionToDecimal(a int, b int) string {
	if a == 0 {
		return "0"
	}
	negative := false
	if a < 0 {
		negative = !negative
		a = -a
	}
	if b < 0 {
		negative = !negative
		b = -b
	}

	ints := "0"
	if a > b {
		ints = strconv.Itoa(a / b)
		a = a % b
	}

	floats := ""
	m := map[int]int{}

	k := 0
	for a != 0 {
		if prev, ok := m[a]; ok {
			floats = floats[:prev] + "(" + floats[prev:] + ")"
			break
		}
		m[a] = k
		a *= 10
		floats += strconv.Itoa(a / b)
		a = a % b
		k += 1
	}
	if floats != "" {
		floats = "." + floats
	}
	if negative {
		return "-" + ints + floats
	}
	return ints + floats
}
func main() {
	fmt.Println(fractionToDecimal(1, 99))
	fmt.Println(fractionToDecimal(1, 3))
}
