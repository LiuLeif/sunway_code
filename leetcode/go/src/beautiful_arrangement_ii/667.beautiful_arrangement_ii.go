// 2018-11-08 12:46
package main

import "fmt"

func constructArray(n int, k int) []int {
	// 1 2 3 4
	ret := []int{}
	lo, hi := 1, n

	flag := true
	for ; k > 0; k-- {
		if flag {
			ret = append(ret, hi)
			hi--
		} else {
			ret = append(ret, lo)
			lo++
		}
		flag = !flag
	}
	if flag {
		for i := lo; i <= hi; i++ {
			ret = append(ret, i)
		}
	} else {
		for i := hi; i >= lo; i-- {
			ret = append(ret, i)
		}
	}
	return ret
}

func main() {
	fmt.Println(constructArray(4, 2))
}
