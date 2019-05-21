// 2017-11-14 18:54
package main

import (
	"fmt"
	"strconv"
)

func dfs(nums []int, k int) string {
	if len(nums) == 0 {
		return ""
	}
	count := 1
	for i := 1; i < len(nums); i++ {
		count *= i
	}

	thisIndex := k/count - 1
	remainder := k % count
	if remainder != 0 {
		thisIndex++
	}
	nextIndex := k - count*(thisIndex)
	return strconv.Itoa(nums[thisIndex]) + dfs(append(nums[:thisIndex], nums[thisIndex+1:]...), nextIndex)
}

func getPermutation(n int, k int) string {
	// 123
	// 132
	// 213
	// 231
	// 312
	// 321
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = i + 1
	}
	return dfs(nums, k)
}

func main() {
	fmt.Println(getPermutation(3, 1))
}
