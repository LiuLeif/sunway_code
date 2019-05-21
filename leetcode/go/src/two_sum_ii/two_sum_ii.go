// 2017-11-28 18:55
package main

import "fmt"

func twoSum(nums []int, target int) []int {
	ret := [2]int{}
	left, right := 0, len(nums)-1
	for left < right {
		if nums[left]+nums[right] == target {
			ret[0] = left + 1
			ret[1] = right + 1
			break
		}
		if nums[left]+nums[right] > target {
			right--
		} else {
			left++
		}
	}

	return ret[:]
}
func main() {
	fmt.Println(twoSum([]int{1, 2, 3, 4}, 4))
}
