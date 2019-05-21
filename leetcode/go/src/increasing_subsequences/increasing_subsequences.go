// 2017-12-25 17:35
package main

import "fmt"

func helper(ret *[][]int, nums, record []int, start int) {
	if len(record) > 1 {
		*ret = append(*ret, append([]int{}, record...))
	}
	if start == len(nums) {
		return
	}
	visited := map[int]bool{}
	for i := start; i < len(nums); i++ {
		if visited[nums[i]] {
			continue
		}
		if len(record) == 0 || nums[i] >= record[len(record)-1] {
			record = append(record, nums[i])
			helper(ret, nums, record, i+1)
			record = record[:len(record)-1]
		}
	}
}

func findSubsequences(nums []int) [][]int {
	ret := [][]int{}
	helper(&ret, nums, []int{}, 0)
	return ret
}
func main() {
	fmt.Println(findSubsequences([]int{1, 2, 3, 4, 5}))
}
