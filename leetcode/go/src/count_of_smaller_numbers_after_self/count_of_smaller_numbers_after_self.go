// 2017-11-14 18:54
package main

import "fmt"

func mergeSort(nums [][]int, ret []int) {
	if len(nums) <= 1 {
		return
	}
	mid := len(nums) / 2
	mergeSort(nums[:mid], ret)
	mergeSort(nums[mid:], ret)

	for i, j := 0, mid; i < mid; i++ {
		for j < len(nums) && nums[j][1] < nums[i][1] {
			j++
		}
		ret[nums[i][0]] += j - mid
	}
	f := 0
	for f < mid && mid < len(nums) {
		if nums[f][1] <= nums[mid][1] {
			f++
		} else {
			t := nums[mid]
			for i := mid; i > f; i-- {
				nums[i] = nums[i-1]
			}
			nums[f] = t
			f++
			mid++
		}
	}
}

func countSmaller(nums []int) []int {
	y := make([][]int, len(nums))
	for k, v := range nums {
		y[k] = []int{k, v}
	}
	ret := make([]int, len(nums))
	mergeSort(y, ret)
	return ret
}

// 2 3, 1 2
// f    mid
func main() {
	fmt.Println(countSmaller([]int{1, 2, 1}))
}
