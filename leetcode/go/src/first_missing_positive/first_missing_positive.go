// 2017-11-14 18:54
package main

import "fmt"

func firstMissingPositive(nums []int) int {
	n := len(nums)
	mark := 0
	max := 0
	for _, v := range nums {
		if v > 0 {
			mark = v
			if v > max {
				max = v
			}
		}
	}
	if mark == 0 {
		return 1
	}

	for i := 0; i < n; i++ {
		if nums[i] <= 0 {
			nums[i] = mark
		}
	}

	fmt.Println(nums)
	for i := 0; i < n; i++ {
		x := nums[i]
		abs := x
		if x < 0 {
			abs = -x
		}

		if abs > n {
			continue
		}
		if nums[abs-1] > 0 {
			nums[abs-1] = -nums[abs-1]
		}
	}

	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			return i + 1
		}
	}

	return max + 1
}
func main() {
	fmt.Println(firstMissingPositive([]int{1, 3, 19}))
}
