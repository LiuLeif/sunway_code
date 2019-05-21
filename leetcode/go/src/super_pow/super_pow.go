// 2018-03-22 13:14
package main

import "fmt"

func pow(a, b int) int {
	ret := 1
	for i := 0; i < b; i++ {
		ret *= a
		ret %= 1337
	}
	return ret
}

func superPow(a int, b []int) int {
	// mod 1337
	if len(b) == 0 {
		return 1
	}
	digit := b[len(b)-1]
	return (pow(superPow(a, b[:len(b)-1]), 10) * pow(a, digit)) % 1337
}

func main() {
	fmt.Println(superPow(2, []int{1, 0, 0, 0, 0, 0, 1}))
}
