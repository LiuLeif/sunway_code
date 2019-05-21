// 2018-02-28 17:41
package main

import (
	"fmt"
	"sort"
)

type Pair struct {
	i, j int
}

func firstLargerEqual(nums []int, x int) int {
	index := sort.Search(len(nums), func(i int) bool {
		return nums[i] >= x
	})
	if index == len(nums) {
		return -1
	}
	return nums[index]
}

func lastSmallerEqual(nums []int, x int) int {
	index := sort.Search(len(nums), func(i int) bool {
		return nums[i] > x
	})
	if index == 0 {
		return -1
	}
	return nums[index-1]
}

func countPalindromicSubsequences(S string) int {
	pos := map[byte][]int{}
	for i := 0; i < len(S); i++ {
		c := S[i]
		if pos[c] == nil {
			pos[c] = []int{}
		}
		pos[c] = append(pos[c], i)
	}
	mem := map[Pair]int{}

	var dfs func(f, t int) int
	dfs = func(f, t int) int {
		if f > t {
			return 0
		}
		if cache, ok := mem[Pair{f, t}]; ok {
			return cache
		}
		ret := 0
		for c := 'a'; c < 'z'; c++ {
			left := firstLargerEqual(pos[byte(c)], f)
			right := lastSmallerEqual(pos[byte(c)], t)
			if left == -1 || left > t {
				continue
			}
			ret++
			if left != right {
				ret++
			}
			ret += dfs(left+1, right-1)
		}
		ret %= 1000000007
		mem[Pair{f, t}] = ret
		return ret
	}

	return dfs(0, len(S)-1)
}
func main() {
	fmt.Println(countPalindromicSubsequences("abcdabcdabcdabcdabcdabcdabcdabcddcbadcbadcbadcbadcbadcbadcbadcba"))
	fmt.Println(countPalindromicSubsequences("axa"))
}
