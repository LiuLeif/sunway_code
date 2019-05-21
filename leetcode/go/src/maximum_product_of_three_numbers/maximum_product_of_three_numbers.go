// 2018-01-10 10:00
package main

import (
	"fmt"
	"math"
	"sort"
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

func maximumProduct(nums []int) int {
	sort.Ints(nums)
	return max(nums[0]*nums[1]*nums[len(nums)-1], nums[len(nums)-1]*nums[len(nums)-2]*nums[len(nums)-3])
}
func main() {
	fmt.Println(maximumProduct([]int{1, 2, 3, 4}))
}
