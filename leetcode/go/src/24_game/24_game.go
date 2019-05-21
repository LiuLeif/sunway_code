// 2018-01-22 16:40
package main

import (
	"fmt"
	"math"
)

func helper(nums []float64) bool {
	if len(nums) == 1 {
		if math.Abs(nums[0]-float64(24)) < 1e-10 {
			return true
		} else {
			return false
		}
	}
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if i == j {
				continue
			}
			nums2 := []float64{}
			for k := 0; k < len(nums); k++ {
				if k != i && k != j {
					nums2 = append(nums2, nums[k])
				}
			}
			// +
			if i < j {
				nums2 = append(nums2, nums[i]+nums[j])
				if helper(nums2) {
					return true
				}
				nums2 = nums2[:len(nums2)-1]
			}
			// -
			nums2 = append(nums2, nums[i]-nums[j])
			if helper(nums2) {
				return true
			}
			nums2 = nums2[:len(nums2)-1]
			// *
			if i < j {
				nums2 = append(nums2, nums[i]*nums[j])
				if helper(nums2) {
					return true
				}
				nums2 = nums2[:len(nums2)-1]
			}
			// /
			if nums[j] != 0 {
				nums2 = append(nums2, nums[i]/nums[j])
				if helper(nums2) {
					return true
				}
				nums2 = nums2[:len(nums2)-1]
			}
		}
	}
	return false
}
func judgePoint24(nums []int) bool {

	floatNums := make([]float64, 4)
	for i := 0; i < 4; i++ {
		floatNums[i] = float64(nums[i])
	}
	return helper(floatNums)
}
func main() {
	fmt.Println(judgePoint24([]int{4, 1, 8, 7}))
}
