// 2018-11-02 16:06
package main

import "fmt"

func uniqueLetterString(S string) int {
	// 8
	// ABA
	// A:2+2
	// B:
	counter := map[byte][]int{}
	for i := 0; i < len(S); i++ {
		if counter[S[i]] == nil {
			counter[S[i]] = []int{}
		}
		counter[S[i]] = append(counter[S[i]], i)
	}
	// ABA
	// -1 0.2 3
	ret := 0
	for _, v := range counter {
		v = append(v, len(S))
		v = append([]int{-1}, v...)
		for i := 0; i < len(v)-2; i++ {
			ret += (v[i+1] - v[i]) * (v[i+2] - v[i+1])
		}
	}
	return ret
}

func main() {
	fmt.Println(uniqueLetterString("ABA"))
}
