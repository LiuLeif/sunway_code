// 2017-11-14 18:54
package main

import "fmt"

func dfs(record []int, visited map[int]bool, nums []int, ret *[][]int) {
	if len(nums) == len(record) {
		tmp := make([]int, len(record))
		copy(tmp, record)
		*ret = append(*ret, tmp)
		return
	}
	for i := 0; i < len(nums); i++ {
		if visited[i] {
			continue
		}
		record = append(record, nums[i])
		visited[i] = true
		dfs(record, visited, nums, ret)
		visited[i] = false
		record = record[:len(record)-1]
	}
}

func permute(nums []int) [][]int {
	ret := [][]int{}
	visited := make(map[int]bool)
	dfs([]int{}, visited, nums, &ret)
	return ret
}
func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}
