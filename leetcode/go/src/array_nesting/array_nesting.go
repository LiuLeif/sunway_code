// 2018-01-03 14:30
package main

import (
	"fmt"
	"math"
)

func max(nums ...int) int {
	ret := math.MinInt32
	for _, n := range nums {
		if ret < n {
			ret = n
		}
	}
	return ret
}

func helper(nums []int, visited map[int]bool, start int) int {
	ret := 0
	for !visited[start] {
		ret++
		visited[start] = true
		start = nums[start]
	}
	return ret
}

func arrayNesting(nums []int) int {
	visited := map[int]bool{}
	ret := 0
	for i := 0; i < len(nums); i++ {
		if !visited[i] {
			ret = max(ret, helper(nums, visited, i))
		}
	}
	return ret
}
func main() {
	fmt.Println(arrayNesting([]int{5, 4, 0, 3, 1, 6, 2}))
}
