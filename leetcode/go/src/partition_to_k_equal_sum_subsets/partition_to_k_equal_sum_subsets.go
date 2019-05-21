// 2017-11-14 18:54
package main

import "fmt"

func partition(nums []int, visited []bool, currentSum int, sum int, k int, start int) bool {
	if currentSum == 0 && k == 0 {
		return true
	}
	ret := false
	for i := start; i < len(nums); i++ {
		if visited[i] == false {
			visited[i] = true
			if currentSum+nums[i] == sum {
				ret = partition(nums, visited, 0, sum, k-1, 0)
				if ret {
					return true
				}
			} else if currentSum+nums[i] < sum {
				ret = partition(nums, visited, currentSum+nums[i], sum, k, start+1)
				if ret {
					return true
				}
			}
			visited[i] = false
		}
	}

	return ret
}

func canPartitionKSubsets(nums []int, k int) bool {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum%k != 0 {
		return false
	}
	sum /= k
	fmt.Println(sum)
	currentSum := 0
	visited := make([]bool, len(nums))
	return partition(nums, visited, currentSum, sum, k, 0)
}
func main() {
	fmt.Println(canPartitionKSubsets([]int{5, 5, 5, 5, 16, 4, 4, 4, 4, 4, 3, 3, 3, 3, 4}, 4))
	fmt.Println(canPartitionKSubsets([]int{605, 454, 322, 218, 8, 19, 651, 2220, 175, 710, 2666, 350, 252, 2264, 327, 1843}, 4))
	fmt.Println(canPartitionKSubsets([]int{4, 4, 4, 4, 1, 1, 1, 5}, 4))
}
