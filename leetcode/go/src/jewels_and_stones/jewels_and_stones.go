// 2018-10-16 09:37
package main

import "fmt"

func numJewelsInStones(J string, S string) int {
	mark := map[rune]bool{}
	for _, c := range J {
		mark[c] = true
	}

	ret := 0
	for _, c := range S {
		if mark[c] {
			ret++
		}
	}
	return ret
}
func main() {
	fmt.Println(numJewelsInStones("aA", "aAAbbbb"))
}
