// 2017-11-14 18:54
package main

import "fmt"

func minSteps(n int) int {
	ret := 0
	for i := 2; i <= n; i++ {
		for n%i == 0 {
			ret += i
			n /= i
		}
	}
	return ret
}
func main() {
	fmt.Println(minSteps(1024))
}
