// 2017-12-01 11:55
package main

import (
	"fmt"
	"math"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	target := 0
	targetLength := math.MaxInt32
	for i := 0; i < len(strs); i++ {
		if len(strs[i]) <= targetLength {
			target = i
			targetLength = len(strs[i])
		}
	}

	targetStr := strs[target]
	for i := 0; i < targetLength; i++ {
		for j := 0; j < len(strs); j++ {
			if targetStr[i] != strs[j][i] {
				return targetStr[:i]
			}
		}
	}

	return targetStr
}
func main() {
	fmt.Println(longestCommonPrefix([]string{"a"}))
}
