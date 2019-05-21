// 2018-01-31 14:37
package main

import "fmt"

func summaryRanges2(nums []int) []string {
	ret := []string{}
	if len(nums) == 0 {
		return ret
	}
	left := 0
	for i := 0; i < len(nums); i++ {
		for i+1 < len(nums) && nums[i+1]-nums[i] == 1 {
			i++
		}
		if left == i {
			ret = append(ret, fmt.Sprintf("%d", nums[left]))
		} else {
			ret = append(ret, fmt.Sprintf("%d->%d", nums[left], nums[i]))
		}
		left = i + 1
	}
	return ret
}

func summaryRanges(nums []int) []string {
	ret := []string{}
	if len(nums) == 0 {
		return ret
	}
	left := 0
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] != nums[i+1]-1 {
			if i == left {
				ret = append(ret, fmt.Sprintf("%d", nums[left]))
			} else {
				ret = append(ret, fmt.Sprintf("%d->%d", nums[left], nums[i]))
			}

			left = i + 1
		}
	}
	if len(nums)-1 == left {
		ret = append(ret, fmt.Sprintf("%d", nums[left]))
	} else {
		ret = append(ret, fmt.Sprintf("%d->%d", nums[left], nums[len(nums)-1]))
	}
	return ret
}
func main() {
	fmt.Println(summaryRanges([]int{1, 2, 3, 4, 5, 8, 9, 11}))
	fmt.Println(summaryRanges2([]int{1, 2, 3, 4, 5, 8, 9, 11}))
}
