// 2017-12-13 13:52
package main

import (
	"fmt"
	"math"
)

func findMaxAverage(nums []int, k int) float64 {
	if len(nums) < k {
		return 0
	}
	sum := 0
	ret := float64(math.MinInt32)
	for i := 0; i < k; i++ {
		sum += nums[i]
	}
	ret = math.Max(ret, float64(sum)/float64(k))
	for i := 1; i < len(nums)-k+1; i++ {
		sum -= nums[i-1]
		sum += nums[i+k-1]
		ret = math.Max(ret, float64(sum)/float64(k))
	}
	return ret
}
func main() {
	fmt.Println(findMaxAverage([]int{1, 12, -5, -6, 50, 3}, 4))
}
