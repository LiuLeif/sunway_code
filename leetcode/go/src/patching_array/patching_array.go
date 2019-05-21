// 2018-03-30 10:26
package main

import "fmt"

func minPatches(nums []int, n int) int {
	// 1,5,10
	// 1,2,4
	check := 0
	sum := 0
	ret := 0
	for sum < n {
		if check >= len(nums) || nums[check] > sum+1 {
			sum += sum + 1
			ret++
		} else {
			sum += nums[check]
			check++
		}
	}
	return ret
}
func main() {
	fmt.Println(minPatches([]int{1, 5, 10}, 20))
	fmt.Println(minPatches([]int{1, 3}, 6))
}
