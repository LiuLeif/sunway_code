// 2017-12-20 17:40
package main

import "fmt"

func convertToTitle(n int) string {
	// 1->A,...26->Z,27->AA,28->AB
	if n <= 26 {
		return string('A' - 1 + n)
	}
	x, y := (n-1)/26, (n-1)%26+1
	return convertToTitle(x) + convertToTitle(y)
}
func main() {
	fmt.Println(convertToTitle(703))
}
