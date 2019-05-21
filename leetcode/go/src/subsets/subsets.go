// 2017-12-06 15:10
package main

import (
	"fmt"
	"sort"
)

func dfs(ret *[][]int, record []int, nums []int, start int) {
	*ret = append(*ret, append([]int{}, record...))
	visited := map[int]bool{}
	for i := start; i < len(nums); i++ {
		if visited[nums[i]] {
			continue
		}
		visited[nums[i]] = true
		record = append(record, nums[i])
		dfs(ret, record, nums, i+1)
		record = record[:len(record)-1]
	}
}

func subsets(nums []int) [][]int {
	ret := [][]int{}
	sort.Ints(nums)
	dfs(&ret, []int{}, nums, 0)
	return ret
}
func main() {
	fmt.Println(subsets([]int{1, 2, 3, 4}))
}
