// 2018-10-25 14:33
package main

import (
	"fmt"
	"math"
	"sort"
)

func numFactoredBinaryTrees(A []int) int {
	MOD := int(math.Pow(10, 9)) + 7
	dp := map[int]int{}
	counter := map[int]bool{}

	factors := map[int]map[int]bool{}

	for _, v := range A {
		counter[v] = true
	}

	nums := []int{}
	for k, _ := range counter {
		nums = append(nums, k)
	}
	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			t := nums[i] * nums[j]
			if counter[t] {
				if factors[t] == nil {
					factors[t] = map[int]bool{}
				}
				factors[t][nums[i]] = true
			}
		}
	}

	ret := 0
	for _, x := range nums {
		n := 0
		for a, _ := range factors[x] {
			b := x / a
			tmp := dp[a] * dp[b]
			if b != a {
				tmp *= 2
			}
			n += tmp
		}
		dp[x] = n + 1
		ret += dp[x]
		ret %= MOD
	}
	return ret
}

func main() {
	fmt.Println(numFactoredBinaryTrees([]int{2, 4, 5, 10}))
}
