// 2017-11-14 18:54
package main

import "fmt"

func missingNumber(nums []int) int {
	n := len(nums) // 0..n
	sum := 0
	for _, v := range nums {
		sum += v
	}

	sum2 := (1 + n) * n / 2
	return sum2 - sum
}

func missingNumber2(nums []int) int {
	n := len(nums)
	left := 0
	right := n

	for left < right {
		mid := left + (right-left)/2
		count := 0
		for i := 0; i < n; i++ {
			if nums[i] <= mid {
				count++
			}
		}
		if count == mid+1 {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

func main() {
	fmt.Println(missingNumber2([]int{0, 1, 4, 2}))
}
