// 2017-11-14 18:54
package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	ret := make([][]int, 0)
	sort.Ints(nums)
	for k, v := range nums {
		if k != 0 && nums[k-1] == v {
			continue
		}
		want := -v
		i := k + 1
		j := len(nums) - 1
		for i < j && j < len(nums) {
			z := nums[i] + nums[j]
			if z == want {
				ret = append(ret, []int{v, nums[i], nums[j]})
				for nums[i] == nums[i+1] && i < j-1 {
					i += 1
				}
				i += 1
				for nums[j] == nums[j-1] && i < j-1 {
					j -= 1
				}
				j -= 1
				continue
			}
			if z > want {
				j -= 1

			} else {
				i += 1
			}
		}

	}
	return ret
}

func main() {
	fmt.Println(threeSum([]int{-2, 0, 0, 2, 2}))
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Println(threeSum([]int{0, 0, 0, 0, 0}))
}
