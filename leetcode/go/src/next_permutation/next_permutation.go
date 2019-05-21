// 2017-11-16 10:34
package main

import "fmt"

func nextPermutation(nums []int) {
	from, to := -1, -1
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i] <= nums[i-1] {
			continue
		}
		from = i - 1
		for j := i; j < len(nums); j++ {
			if nums[j] > nums[i-1] {
				to = j
			}
		}
		break
	}
	if from != -1 {
		nums[from], nums[to] = nums[to], nums[from]
	}

	// reverse [from+1:]
	tmp := nums[from+1:]
	for i := 0; i < len(tmp)/2; i++ {
		tmp[i], tmp[len(tmp)-i-1] = tmp[len(tmp)-i-1], tmp[i]
	}

}
func main() {
	x := []int{5, 1, 1}
	nextPermutation(x)
	fmt.Println(x)
}
