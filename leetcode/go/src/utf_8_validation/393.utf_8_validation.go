// 2018-11-09 15:12
package main

import "fmt"

func validUtf8(data []int) bool {
	getLeadingOnes := func(x int) int {
		ret := 0
		for i := 7; i >= 0 && x&(1<<uint(i)) != 0; i-- {
			ret++
		}
		return ret
	}
	N := len(data)
	prev := 0
	for i := 0; i < N; i++ {
		leadingOnes := getLeadingOnes(data[i])
		// fmt.Println("leading", data[i], leadingOnes)
		switch leadingOnes {
		case 0:
			if prev != 0 {
				return false
			}
		case 1:
			if prev == 0 {
				return false
			}
			prev--
		default:
			if prev != 0 {
				return false
			}
			prev = leadingOnes - 1
			if prev > 3 {
				return false
			}
		}
	}
	return prev == 0
}

func main() {
	fmt.Println(validUtf8([]int{250, 145, 145, 145, 145}))
}
