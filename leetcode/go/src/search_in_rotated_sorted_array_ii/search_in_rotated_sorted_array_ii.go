// 2018-01-09 15:41
package main

import "fmt"

func search(nums []int, target int) bool {
	f, t := 0, len(nums)-1
	for f <= t {
		mid := f + (t-f)/2
		if nums[mid] == target {
			return true
		}
		if nums[mid] > nums[t] {
			if nums[f] <= target && nums[mid] >= target {
				t = mid - 1
			} else {
				f = mid + 1
			}
		} else if nums[mid] < nums[t] {
			if nums[mid] <= target && nums[t] >= target {
				f = mid + 1
			} else {
				t = mid - 1
			}
		} else {
			t--
		}
	}

	return false
}

func main() {
	fmt.Println(search([]int{1, 3, 5, 1}, 2))
}
