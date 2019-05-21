// 2017-11-14 18:54
package main

import "fmt"

func subarraySum(nums []int, k int) int {
	total := 0
	for _, v := range nums {
		total += v
	}
	count := 0
	currTotal := total
	for i := 0; i < len(nums); i++ {
		for j := len(nums) - 1; j >= i; j-- {
			if currTotal == k {
				count++
			}
			currTotal -= nums[j]
		}
		total -= nums[i]
		currTotal = total
	}

	return count
}
func main() {
	fmt.Println(subarraySum([]int{1, 1, 1}, 0))
}
