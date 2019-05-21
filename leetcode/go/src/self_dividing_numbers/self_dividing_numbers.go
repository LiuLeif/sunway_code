// 2018-01-08 10:37
package main

import "fmt"

func selfDividingNumbers(left int, right int) []int {
	ret := []int{}
	for n := left; n <= right; n++ {
		tmp := n
		for tmp != 0 {
			mod := tmp % 10
			if mod == 0 {
				break
			}
			if n%mod != 0 {
				break
			}
			tmp = tmp / 10
		}
		if tmp == 0 {
			ret = append(ret, n)
		}
	}
	return ret
}
func main() {
	// 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22
	fmt.Println(selfDividingNumbers(1, 10))
}
