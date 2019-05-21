// 2017-11-27 15:46
package main

import "fmt"

func dfs(nums []int, visited []bool, sum, currentSum, k, start int) bool {
	if k == 0 {
		return true
	}
	if currentSum == sum {
		return dfs(nums, visited, sum, 0, k-1, 0)
	}
	for i := start; i < len(nums); i++ {
		if visited[i] {
			continue
		}
		visited[i] = true
		if currentSum+nums[i] <= sum {
			ret := dfs(nums, visited, sum, currentSum+nums[i], k, i+1)
			if ret {
				return true
			}
		}
		visited[i] = false
	}

	return false
}

func makesquare(nums []int) bool {
	if len(nums) == 0 {
		return false
	}
	sum := 0
	for _, v := range nums {
		sum += v
	}

	if sum%4 != 0 {
		return false
	}
	sum /= 4
	visited := make([]bool, len(nums))
	return dfs(nums, visited, sum, 0, 4, 0)
}
func main() {
	fmt.Println(makesquare([]int{5, 5, 5, 5, 16, 4, 4, 4, 4, 4, 3, 3, 3, 3, 4}))
}
