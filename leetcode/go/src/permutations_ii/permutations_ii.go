// 2017-11-14 18:54
package main

import "fmt"

func dfs(record []int, visited map[int]bool, nums []int, ret *[][]int) {
	if len(record) == len(nums) {
		t := make([]int, len(record))
		copy(t, record)
		*ret = append(*ret, t)
		return
	}
	dup := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		if visited[i] {
			continue
		}
		if dup[nums[i]] {
			continue
		}
		dup[nums[i]] = true

		visited[i] = true
		record = append(record, nums[i])
		dfs(record, visited, nums, ret)
		record = record[:len(record)-1]
		visited[i] = false
	}
}

func permuteUnique(nums []int) [][]int {
	ret := [][]int{}
	visited := make(map[int]bool)
	dfs([]int{}, visited, nums, &ret)
	return ret
}
func main() {
	fmt.Println(permuteUnique([]int{1, 1, 2}))
}
