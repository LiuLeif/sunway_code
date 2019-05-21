// 2018-11-02 14:41
package main

import (
	"fmt"
	"sort"
)

func orderlyQueue(S string, K int) string {
	if K != 1 {
		v := []rune(S)
		sort.Slice(v, func(i, j int) bool {
			return v[i] < v[j]
		})
		return string(v)
	}
	minS := S
	for i := 0; i < len(S); i++ {
		t := S[i:] + S[:i]
		if t < minS {
			minS = t
		}
	}
	return minS
}

func main() {
	fmt.Println(orderlyQueue("baaca", 1))
}
