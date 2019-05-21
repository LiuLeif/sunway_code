// 2017-12-05 16:02
package main

import (
	"fmt"
	"math"
)

func checkPerfectNumber(num int) bool {
	if num == 1 {
		return false
	}
	sum := 1
	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			sum += i
			if num/i != i {
				sum += num / i
			}
		}
	}
	return num == sum
}
func main() {
	fmt.Println(checkPerfectNumber(28))
}
