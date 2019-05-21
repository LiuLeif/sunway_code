// 2017-11-14 18:54
package main

import "fmt"

func max(nums ...int) int {
	ret := 0
	for _, n := range nums {
		if ret < n {
			ret = n
		}
	}
	return ret
}

func longestPalindromeSubseq(s string) int {
	m := make([][]int, len(s))
	for i := 0; i < len(s); i++ {
		m[i] = make([]int, len(s))
	}

	for i := 0; i < len(s); i++ {
		m[i][i] = 1
	}

	ret := 1
	for i := 1; i < len(s); i++ {
		for j := 0; j < len(s)-i; j++ {
			if i == 1 {
				if s[j] == s[j+i] {
					m[j][j+i] = 2
				} else {
					m[j][j+i] = 1
				}
			} else {
				if s[j] == s[j+i] {
					m[j][j+i] = max(m[j+1][j+i-1]+2, m[j][j+i-1], m[j+1][j+i])
				} else {
					m[j][j+i] = max(m[j][j+i-1], m[j+1][j+i])
				}
			}
			ret = max(ret, m[j][j+i])
		}
	}

	return ret
}

func main() {
	fmt.Println(longestPalindromeSubseq("aabxa"))
}
