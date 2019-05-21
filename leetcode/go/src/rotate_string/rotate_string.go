// 2018-03-30 10:13
package main

import (
	"fmt"
	"strings"
)

func rotateString(A string, B string) bool {
	return len(A) == len(B) && strings.Index(A+A, B) != -1
}
func main() {
	fmt.Println(rotateString("abcde", "cdeab"))
	fmt.Println(rotateString("", ""))
}
