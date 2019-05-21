// 2018-02-01 11:35
package main

import "fmt"

func arrangeCoins(n int) int {
	ret := 0
	sum := 0
	step := 1
	for sum <= n {
		sum += step
		step++
		ret++
	}
	return ret - 1
}
func main() {
	fmt.Println(arrangeCoins(5))
}
