// 2018-11-13 17:51
package main

import (
	"fmt"
	"math"
)

func findMinStep(board string, hand string) int {
	min := func(nums ...int) int {
		ret := math.MaxInt32
		for _, n := range nums {
			if ret > n {
				ret = n
			}
		}
		return ret
	}
	memo := map[string]int{}
	reduce := func(S string) string {
		ret := ""
		for {
			ret = ""
			// aaabbbc
			// i j
			i := 0
			for j := 0; j < len(S); j++ {
				if S[i] != S[j] {
					if j-i < 3 {
						ret += S[i:j]
					}
					i = j
				}
			}
			if len(S)-i < 3 {
				ret += S[i:]
			}

			if S == ret {
				break
			}
			S = ret
		}
		return ret
	}

	var dfs func(S, s string) int
	dfs = func(S, s string) int {
		if len(S) == 0 {
			return 0
		}
		if len(s) == 0 {
			return math.MaxInt32
		}
		key := S + "," + s
		if memo[key] != 0 {
			return memo[key]
		}

		ret := math.MaxInt32

		for j := 0; j < len(s); j++ {
			c := s[j]
			s2 := s[:j] + s[j+1:]
			for i := 0; i < len(S); i++ {
				if i != 0 && S[i] == S[i-1] {
					continue
				}
				if S[i] != c {
					continue
				}
				S2 := reduce(S[:i] + string(c) + S[i:])
				ret = min(ret, 1+dfs(S2, s2))
			}
		}

		memo[key] = ret
		return ret
	}
	ret := dfs(board, hand)
	if ret > len(hand) {
		return -1
	}
	return ret
}

func main() {
	fmt.Println(findMinStep("WRYYRWWRRWW", "WYBR"))
	fmt.Println(findMinStep("WWRRBBWW", "WRBRW"))
	fmt.Println(findMinStep("G", "GGGGG"))
	fmt.Println(findMinStep("RBYYBBRRB", "YRBGB"))
	fmt.Println(findMinStep("WGBBGYGWGWWRYB", "GWGRW"))
}
