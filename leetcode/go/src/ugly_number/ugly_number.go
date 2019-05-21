// 2017-12-29 16:35
package main

import "fmt"

func isUgly(num int) bool {
	if num <= 0 {
		return false
	}
	// 2,3,5
	primes := []int{2, 3, 5}
	for _, prime := range primes {
		for num%prime == 0 {
			num /= prime
		}
	}
	return num == 1
}
func main() {
	fmt.Println(isUgly(0))
}
