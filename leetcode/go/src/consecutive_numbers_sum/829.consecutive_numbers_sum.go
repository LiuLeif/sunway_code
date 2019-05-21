// 2018-11-05 15:45
package main

import "fmt"

func consecutiveNumbersSum(N int) int {
	ret := 0
	for i := 1; (i+1)*i/2 <= N; i++ {
		if (N-(i+1)*i/2)%i == 0 {
			ret++
		}
	}
	return ret
}

func main() {
	fmt.Println(consecutiveNumbersSum(15))
}
