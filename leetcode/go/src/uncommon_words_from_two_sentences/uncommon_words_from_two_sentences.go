// 2018-10-12 15:34
package main

import "fmt"
import "strings"

func uncommonFromSentences(A string, B string) []string {
	ma := map[string]int{}
	mb := map[string]int{}
	for _, v := range strings.Split(A, " ") {
		ma[v] += 1
	}

	for _, v := range strings.Split(B, " ") {
		mb[v] += 1
	}

	ret := []string{}
	for k, v := range ma {
		if v == 1 && mb[k] == 0 {
			ret = append(ret, k)
		}
	}
	for k, v := range mb {
		if v == 1 && ma[k] == 0 {
			ret = append(ret, k)
		}
	}
	return ret
}

func main() {
	fmt.Println(uncommonFromSentences("apple apple", "banana banana"))
}
