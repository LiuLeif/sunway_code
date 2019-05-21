// 2018-01-26 11:11
package main

import "fmt"

func helper(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	mid := n / 2
	mod := n % 2
	t := myPow(x, mid)
	if mod == 0 {
		return t * t
	} else {
		return t * t * x
	}
}

func myPow(x float64, n int) float64 {
	if n >= 0 {
		return helper(x, n)
	} else {
		return 1 / helper(x, -n)
	}

}
func main() {
	// fmt.Println(myPow(2, -2))
	fmt.Println(myPow(0.00001, 2147483647))
}
