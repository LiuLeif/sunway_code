// 2018-01-24 13:32
package main

import (
	"fmt"
	"strconv"
)

func fizzBuzz(n int) []string {
	ret := make([]string, n)
	for i := 1; i <= n/3; i++ {
		ret[i*3-1] += "Fizz"
	}
	for i := 1; i <= n/5; i++ {
		ret[i*5-1] += "Buzz"
	}
	for i := 0; i < n; i++ {
		if len(ret[i]) == 0 {
			ret[i] = strconv.Itoa(i + 1)
		}
	}
	return ret
}
func main() {
	fmt.Println(fizzBuzz(16))
}
