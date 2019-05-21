// 2018-11-15 14:00
package main

import (
	"fmt"
	"sort"
	"strings"
)

func reorderLogFiles(logs []string) []string {
	alpha := []string{}
	digit := []string{}
	for _, s := range logs {
		if s[len(s)-1] >= '0' && s[len(s)-1] <= '9' {
			digit = append(digit, s)
		} else {
			alpha = append(alpha, s)
		}
	}
	sort.Slice(alpha, func(i, j int) bool {
		a := strings.Index(alpha[i], " ")
		b := strings.Index(alpha[j], " ")
		if alpha[i][a:] < alpha[j][b:] {
			return true
		}
		if alpha[i][a:] > alpha[j][b:] {
			return false
		}
		return alpha[i] < alpha[j]
	})

	return append(alpha, digit...)
}

func main() {
	fmt.Println(reorderLogFiles([]string{
		"a1 9 2 3 1", "g1 act car", "zo4 4 7", "ab1 off key dog", "a8 act zoo",
	}))
}
