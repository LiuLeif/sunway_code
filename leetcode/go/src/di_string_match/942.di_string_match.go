// 2018-11-19 16:14
package main

import "fmt"

func diStringMatch(S string) []int {
	ret := []int{}
	//   I D I D
	// 0 4 1 3 2
	lo, hi := 0, len(S)
	for i := 0; i < len(S); i++ {
		if S[i] == 'I' {
			ret = append(ret, lo)
			lo++
		} else {
			ret = append(ret, hi)
			hi--
		}
	}
	ret = append(ret, lo)
	return ret
}

func main() {
	fmt.Println(diStringMatch("DDDII"))
}
