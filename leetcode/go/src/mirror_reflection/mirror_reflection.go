// 2018-10-18 10:41
package main

import "fmt"

func mirrorReflection(p int, q int) int {
	Q := 0
	for i := 0; ; i++ {
		Q += q
		Q %= 2 * p
		if Q == 0 {
			return 0
		}
		if Q == p {
			if i%2 == 1 {
				return 2
			} else {
				return 1
			}
		}
	}
	return 0
}
func main() {
	fmt.Println(mirrorReflection(2, 1))
}
