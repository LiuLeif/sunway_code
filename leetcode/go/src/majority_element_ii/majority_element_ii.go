// 2017-12-05 10:50
package main

import (
	"fmt"
	"math"
)

func majorityElement(nums []int) []int {
	candidate1, candidate2 := 0, 0
	count1, count2 := 0, 0

	for _, v := range nums {
		if count1 != 0 && v == candidate1 {
			count1++
		} else if count2 != 0 && v == candidate2 {
			count2++
		} else if count1 == 0 {
			count1++
			candidate1 = v
		} else if count2 == 0 {
			count2++
			candidate2 = v
		} else {
			count1--
			count2--
		}
	}
	ret := []int{}
	count1, count2 = 0, 0
	for _, v := range nums {
		if v == candidate1 {
			count1++
		} else if v == candidate2 {
			count2++
		}
	}

	if count1 > int(math.Floor(float64(len(nums))/3.0)) {
		ret = append(ret, candidate1)
	}
	if count2 > int(math.Floor(float64(len(nums))/3.0)) {
		ret = append(ret, candidate2)
	}
	return ret
}
func main() {
	fmt.Println(majorityElement([]int{1}))
}
