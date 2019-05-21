// 2018-04-13 11:07
package main

import "fmt"

func dfs(nums []int) int {
	if len(nums) == 0 {
		return 1
	}
	N := len(nums)
	ret := 0
	for i := 0; i < len(nums); i++ {
		if nums[i]%N == 0 || N%nums[i] == 0 {
			nums[i], nums[N-1] = nums[N-1], nums[i]
			ret += dfs(nums[0 : N-1])
			nums[i], nums[N-1] = nums[N-1], nums[i]
		}
	}
	return ret
}

func countArrangement(N int) int {
	// 1 2 3 4 5
	nums := make([]int, N)
	for i := 0; i < N; i++ {
		nums[i] = i + 1
	}
	return dfs(nums)
}
func main() {
	fmt.Println(countArrangement(3))
}
